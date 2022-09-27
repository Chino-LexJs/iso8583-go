package prosa

import (
	"fmt"
	"net"
	"os"
)

const (
	PROSA_HOST = "localhost"
	PROSA_PORT = "8000"
)

func SendMessageProsa(message string, done chan string) {
	done <- connectProsa(message)
}

func connectProsa(message string) string {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error al inicar conexion con PROSA:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
	} else {
		fmt.Println("\nMessage send ok")
	}
	buffer := make([]byte, 1024)
	for {
		mLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		if mLen != 0 {
			fmt.Println("\nReceived: ", string(buffer[:mLen]))
			return string(buffer[:mLen])
		}
	}
	return ""
}
