package modules

import (
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

func (client RTSP) IsConnected() bool {
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
	for {
		if retry < 10 && !client.IsConnected() {
			socket, err := net.DialTimeout("tcp", net.JoinHostPort(client.Ip, client.Port), time.Duration(time.Second*10))
			if err != nil {
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

}

func (client *RTSP) Describe() []byte {
	auth_str := ""
	result := fmt.Sprintf("DESCRIBE rtsp://%s:%s%s RTSP/1.0\r\nCSeq: %d\r\n%sUser-Agent: Mozilla/5.0\r\nAccept: application/sdp\r\n\r\n", client.Ip, client.Port, "/1", client.Cseq, auth_str)
	fmt.Println(result)
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

func (client *RTSP) Authorize() bool {
	if client.IsConnected() {
		ReadBuf := make([]byte, 1024)
		_, err := client.Socket.Write(client.Describe())
		if err != nil {
			fmt.Println(err)
		}
		client.Socket.Read(ReadBuf)
		client.Data = string(ReadBuf)

		return true
	} else {
		return false
	}
}
