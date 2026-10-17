package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	buffer := []byte("Hello this is Animesh")
	n, err := conn.Write(buffer)

	if err != nil {
		fmt.Println("Failed to send buffer: ", err)
		return
	}
	fmt.Println("Wrote buffer of length: ", n)

	readBuff := make([]byte, 1024)
	cn, err := conn.Read(readBuff)

	if err != nil {
		fmt.Println("Failed to read buffer: ", err)
		return
	}
	fmt.Println("Total length of echo buffer read: ", cn)
	fmt.Println("Echo buffer: ", string(readBuff[:n]))
}
