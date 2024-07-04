package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
)

func main() {
	// Parse command line arguments
	ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
	port := flag.Int("port", 0, "Port to bind to")
	flag.Parse()

	// If port is not specified, find the first available port between 80 and 60000
	if *port == 0 {
		*port = findAvailablePort(*ip)
		if *port == 0 {
			fmt.Println("No available port found between 80 and 60000")
			os.Exit(1)
		}
	}

	addr := fmt.Sprintf("%s:%d", *ip, *port)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Printf("Error listening on %s: %v\n", addr, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Server started on %s\n", addr)

	numWorkers := runtime.NumCPU()
	fmt.Printf("Using %d workers\n", numWorkers)

	var wg sync.WaitGroup
	requests := make(chan request, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(conn, requests, &wg)
	}

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Printf("Error reading from connection: %v\n", err)
			continue
		}

		req := request{
			data:       append([]byte{}, buffer[:n]...),
			clientAddr: clientAddr,
		}
		requests <- req
	}

	close(requests)
	wg.Wait()
}

type request struct {
	data       []byte
	clientAddr net.Addr
}

// worker handles requests from the requests channel
func worker(conn net.PacketConn, requests chan request, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		// Get client IP and port
		clientIP, clientPort, err := net.SplitHostPort(req.clientAddr.String())
		if err != nil {
			fmt.Printf("Error splitting client address: %v\n", err)
			continue
		}

		// Reply to client with their IP and port in the format ipv4:port
		response := fmt.Sprintf("%s:%s", clientIP, clientPort)
		_, err = conn.WriteTo([]byte(response), req.clientAddr)
		if err != nil {
			fmt.Printf("Error writing to connection: %v\n", err)
			continue
		}

		fmt.Printf("Sent response to %s:%s\n", clientIP, clientPort)
	}
}

// findAvailablePort finds the first available port between 80 and 60000 for the specified IP address
func findAvailablePort(ip string) int {
	for port := 80; port <= 60000; port++ {
		addr := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.ListenPacket("udp", addr)
		if err == nil {
			conn.Close()
			return port
		}
	}
	return 0
}
