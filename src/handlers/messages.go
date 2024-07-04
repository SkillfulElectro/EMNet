package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"EMNet/server"
	"net"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func addPeer(conn net.Conn, ipPeer string, response map[string]interface{}) {
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	for {
		randomNumber := rand.Intn(1000000) // Generate a random number
		if _, exists := server.Peers_IP[randomNumber]; !exists {
			server.Peers_IP[randomNumber] = ipPeer
			response["msg"] = "SUCCESS"
			response["id"] = randomNumber
			sendResponse(conn, response)
			return
		}
	}
}

func askIP(conn net.Conn, id int, response map[string]interface{}) {
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	if ip, exists := server.Peers_IP[id]; exists {
		response["msg"] = "SUCCESS"
		response["IP"] = ip
		delete(server.Peers_IP, id)
		sendResponse(conn, response)
	} else {
		response["msg"] = "FAILED"
		sendResponse(conn, response)
	}
}

func sendResponse(conn net.Conn, response map[string]interface{}) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	conn.Write(jsonResponse)
	conn.Write([]byte("\n")) // Ensure the message is newline-terminated
}
