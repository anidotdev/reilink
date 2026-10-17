package main

import (
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)

		if n > 0 {
			fmt.Println("Received data:", string(buffer[:n]))

			cn, writeErr := conn.Write(buffer[:n])

			if writeErr != nil {
				fmt.Println("Error echoing data:", writeErr)
				return
			}

			fmt.Println("Sent bytes:", cn)
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected.")
			} else {
				fmt.Println("Error reading data:", err)
			}
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("Failed to bind to port 8080: %v", err)
		return
	}
	defer listener.Close()

	fmt.Println("Reilink is listening on port 8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting the connection: ", err)
			continue
		}
		go handleConnection(conn)
	}

}
