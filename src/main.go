package main

import (
	"flag"
	"fmt"
	"EMNet/server"
	"net"
	"os"
	"runtime"
	"strconv"
)

func main() {
	ip := flag.String("ip", "", "IP address to bind the server")
	port := flag.Int("port", 0, "Port to bind the server")
	flag.Parse()

	if *ip == "" {
		*ip = "0.0.0.0"
	}

	numCores := runtime.NumCPU()
	workerPool := server.NewWorkerPool(numCores) // Create a pool with as many workers as CPU cores

	if *port != 0 {
		// Case 1: -ip and -port are specified
		startServer(*ip, *port, workerPool)
	} else if *ip != "" && *port == 0 {
		// Case 2: only -ip is specified
		for port := 80; port <= 60000; port++ {
			if tryPort(*ip, port) {
				startServer(*ip, port, workerPool)
				break
			}
		}
	} else {
		// Case 3: no CLI args
		for port := 80; port <= 60000; port++ {
			if tryPort("0.0.0.0", port) {
				startServer("0.0.0.0", port, workerPool)
				break
			}
		}
	}
}

func tryPort(ip string, port int) bool {
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func startServer(ip string, port int, workerPool *server.WorkerPool) {
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	fmt.Printf("Server started on %s\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		job := server.Job{Conn: conn}
		workerPool.JobQueue <- job // Send the job to the worker pool
	}
}
