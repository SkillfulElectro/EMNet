package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"EMNet/server"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn) {
	fmt.Printf("Client connected: %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('}')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		if strings.HasPrefix(message, "{") {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(message), &data)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				continue
			}
			processMessage(conn, data)
		} else {
			fmt.Println("Received invalid message format")
		}
	}
}

func processMessage(conn net.Conn, data map[string]interface{}) {
	response := make(map[string]interface{})

	switch data["msg"] {
	case "ADD_PEER":
		ipPeer, ok := data["IP_Peer"].(string)
		if !ok {
			response["msg"] = "FAILED"
			sendResponse(conn, response)
			return
		}
		addPeer(conn, ipPeer, response)
	case "ASK_IP":
		id, ok := data["id"].(float64)
		if !ok {
			response["msg"] = "FAILED"
			sendResponse(conn, response)
			return
		}
		askIP(conn, int(id), response)
	default:
		response["msg"] = "FAILED"
		sendResponse(conn, response)
	}
}
