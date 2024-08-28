package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Availability bool   `json:"availability"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	RandNumber   int    `json:"rand_number"`
}

func main() {
	var res Response

	rand.Seed(time.Now().UnixNano())

	udpServer, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer udpServer.Close()

	buffer := make([]byte, 1024)
	_, addr, err := udpServer.ReadFrom(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res.Availability = true
	res.Host = "localhost"
	res.Port = ":5043"
	res.RandNumber = rand.Intn(4) + 3

	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = udpServer.WriteTo(data, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	listen, err := net.Listen("tcp", res.Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer listen.Close()

	conn, err := listen.Accept()
	buffer = make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
		return
	}

	receivedMessage := string(buffer[:n])
	fmt.Printf("Received message: %v\n", receivedMessage)

	receivedNumber, err := strconv.Atoi(receivedMessage)
	if err != nil {
		fmt.Printf("Error converting message to int: %v\n", err)
		return
	}

	// var recievedInt int32
	// binary.Read(new.NewReader(num_buffer), binary.BigEndian, &recievedInt)

	if receivedNumber != res.RandNumber {
		fmt.Println("Los números no coinciden. Conexión rechazada")
		os.Exit(1)
	}

	fmt.Println("Conexión aceptada")

	for {
		_, err = conn.Write([]byte("pene"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
