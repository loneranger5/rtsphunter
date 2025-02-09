package main

import (
	"fmt"
	"rtspranger/src/modules"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	auth := modules.AUTH{
		Username: "admin",
		Password: "pass",
		Method:   "basic",
	}

	client := modules.CreateRTSPClient("61.43.141.99", "554", 10, auth)
	fmt.Println(client, client.IsConnected())
	if client.Connect("554") && client.Socket != nil {
		fmt.Println(client.Status, client.IsConnected())

		fmt.Println(client.Data)

		client.Authorize()

		client.Socket.Close()

		if len(client.Data) > 0 {
			if strings.Contains(client.Data, "Basic") {
				client.Auth.Method = "Basic"
			} else if strings.Contains(client.Data, "Digest") {
				client.Auth.Method = "Digest"
			}
		}

		fmt.Println(client.Data)
		fmt.Println(client.Auth.Method)

		realm := client.FindRealm()
		nonce := client.FindNonce()

		fmt.Println(realm, nonce)

	}

}
