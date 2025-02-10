package modules

import (
	"crypto/md5"
	b64 "encoding/base64"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

type AUTH struct {
	Username string
	Password string
	Method   string
}
type RTSP struct {
	Ip      string
	Port    string
	Socket  net.Conn
	Status  int
	Realm   string
	Nonce   string
	Route   string
	Timeout int
	Cseq    int
	Data    string
	Auth    AUTH
}

var ROUTE_OK_CODES []string = []string{
	"RTSP/1.0 200",
	"RTSP/1.0 401",
	"RTSP/1.0 403",
	"RTSP/2.0 200",
	"RTSP/2.0 401",
	"RTSP/2.0 403",
}

var CREDENTIALS_OK_CODES []string = []string{"RTSP/1.0 200", "RTSP/1.0 404", "RTSP/2.0 200", "RTSP/2.0 404"}

func MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func base64Encode(text string) string {

	encoded := b64.StdEncoding.EncodeToString([]byte(text))
	return encoded
}
func (client *RTSP) IsConnected() bool {
	if client.Status == 1 {
		return true
	} else {
		return false
	}
}

func CreateRTSPClient(ip string, port string, timeout int, Auth AUTH) RTSP {
	rtsp_client := RTSP{Ip: ip, Port: port, Timeout: timeout, Auth: Auth,
		Nonce: "", Cseq: 0, Data: "", Status: 0}
	return rtsp_client

}

func (client *RTSP) Connect(port string) bool {
	if client.IsConnected() {
		return true
	}

	if port == "0" {
		port = client.Port
	}

	client.Cseq = 0
	client.Data = ""
	retry := 0
	if !client.IsConnected() {
		socket, err := net.DialTimeout("tcp", net.JoinHostPort(client.Ip, client.Port), time.Duration(time.Second*time.Duration(client.Timeout)))
		if err != nil {
			fmt.Println(err)
			retry += 1
			time.Sleep(time.Millisecond)
		}

		if socket != nil {
			client.Socket = socket

		}
		client.Status = 1
		return true

	} else {
		return false
	}

}

func (client *RTSP) Describe(auth_str string) []byte {

	result := fmt.Sprintf("DESCRIBE rtsp://%s:%s%s RTSP/1.0\r\nCSeq: %d\r\n%sUser-Agent: Mozilla/5.0\r\nAccept: application/sdp\r\n\r\n", client.Ip, client.Port, "/", client.Cseq, auth_str)
	return []byte(result)
}

func (client *RTSP) FindRealm() string {
	if len(client.Data) > 0 {
		r, _ := regexp.Compile("realm=\"[a-z0-9]+\"")
		realm := r.Find([]byte(client.Data))
		text := string(realm)
		text = strings.ReplaceAll(text, "\"", "")
		text = strings.ReplaceAll(text, "realm=", "")
		return string(text)
	}

	return ""
}

func (client *RTSP) FindNonce() string {
	if len(client.Data) > 0 {
		r, _ := regexp.Compile("nonce=\"[a-zA-Z0-9]+\"")
		nonce := r.Find([]byte(client.Data))
		text := string(nonce)
		text = strings.ReplaceAll(text, "\"", "")
		text = strings.ReplaceAll(text, "nonce=", "")
		return string(text)

	}

	return ""
}

func (client *RTSP) CheckAuth() {
	if client.IsConnected() {
		if len(client.Data) > 0 {
			if strings.Contains(client.Data, "Basic") {
				client.Auth.Method = "Basic"
			} else if strings.Contains(client.Data, "Digest") {
				client.Auth.Method = "Digest"
			}
		}
	}
}

func (client *RTSP) DigestAuth(option string, route string) string {
	realm, nonce := client.Realm, client.Nonce
	if client.IsConnected() {
		username, password := client.Auth.Username, client.Auth.Password
		uri := fmt.Sprintf("rtsp://%s:%s%s", client.Ip, client.Port, route)
		hash1 := MD5(fmt.Sprintf("%s:%s:%s", username, realm, password))
		hash2 := MD5(fmt.Sprintf("%s:%s", option, uri))
		response := MD5(fmt.Sprintf("%s:%s:%s", hash1, nonce, hash2))

		result := fmt.Sprintf("Authorization: Digest username=\"%s\", realm=\"%s\", nonce=\"%s\", uri=\"%s\", response=\"%s\"", username, realm, nonce, uri, response)

		return result
	} else {
		return ""
	}
}

func (client *RTSP) BasicAuth(credentials string) string {
	encoded := base64Encode(credentials)
	return fmt.Sprintf("Authorization: Basic %s", encoded)
}

func (client *RTSP) OkAuth() bool {
	if client.IsConnected() {
		for _, code := range CREDENTIALS_OK_CODES {
			if strings.Contains(client.Data, code) {
				return true
			}
		}
	}
	return false
}
func (client *RTSP) OkRoute(route string) bool {
	if client.IsConnected() {
		for _, code := range ROUTE_OK_CODES {
			if strings.Contains(client.Data, code) {
				client.Route = route
				return true
			}
		}
	}
	return false
}

func (client *RTSP) GetRTSPUrl() string {
	var url string
	if client.IsConnected() && client.Route != "" {
		if client.Auth.Username == "" && client.Auth.Password == "" {
			url = fmt.Sprintf("rtsp://%s:%s%s", client.Ip, client.Port, client.Route)

		} else {
			url = fmt.Sprintf("rtsp://%s:%s@%s:%s%s", client.Auth.Username, client.Auth.Password, client.Ip, client.Port, client.Route)

		}
		return url
	}
	return ""
}
func (client *RTSP) Authorize(creds string, route string) bool {
	if client.IsConnected() {
		auth_str := ""

		client.Cseq += 1
		ReadBuf := make([]byte, 1024)
		_, err := client.Socket.Write(client.Describe(auth_str))
		if err != nil {
			fmt.Println(err)
		}
		client.Socket.Read(ReadBuf)
		client.Data = string(ReadBuf)
		client.CheckAuth()

		client.Realm = client.FindRealm()
		client.Nonce = client.FindNonce()

		if client.Auth.Method == "Basic" {
			creds := client.Auth.Username + ":" + client.Auth.Password
			auth_str = client.BasicAuth(creds)

		} else if client.Auth.Method == "Digest" {
			auth_str = client.DigestAuth("DESCRIBE", route)
		}

		_, err = client.Socket.Write(client.Describe(auth_str))
		if err != nil {
			fmt.Println(err)
		}
		client.Socket.Read(ReadBuf)

		return true
	} else {
		return false
	}
}
