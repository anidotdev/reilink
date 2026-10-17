package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')

		if msg != "" {
			fmt.Println(msg)
			_, writeErr := conn.Write([]byte(msg))
			if writeErr != nil {
				fmt.Println("Error echoing data:", writeErr)
				return
			}
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

func writeFrame(conn net.Conn, payload []byte) error {
	length := uint32(len(payload))
	header := make([]byte, 4)

	binary.BigEndian.PutUint32(header, length)

	_, err := conn.Write(header)
	if err != nil {
		return err
	}

	_, err = conn.Write(payload)
	if err != nil {
		return err
	}

	return nil
}

func readFrame(conn net.Conn) ([]byte, error) {
	header := make([]byte, 4)

	_, err := io.ReadFull(conn, header)

	if err != nil {
		fmt.Println("Error in reading payload: ", err)
		return nil, err
	}

	length := binary.BigEndian.Uint32(header)
	const maxFrameSize = 1 << 20

	if length > maxFrameSize {
		return nil, fmt.Errorf("frame too large: %d bytes", length)
	}

	payload := make([]byte, int(length))

	_, err = io.ReadFull(conn, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
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
