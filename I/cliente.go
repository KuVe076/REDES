package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

type Response struct {
	Availability bool   `json:"availability"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	RandNumber   int    `json:"rand_number"`
}

func main() {

	strUdp := ":8080"

	str, err := net.ResolveUDPAddr("udp", strUdp)

	conn, err := net.DialUDP("udp", nil, str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	message := []byte("Hola conchetumare")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)

	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res Response
	err = json.Unmarshal(buffer[:n], &res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Disponible: ", res.Availability)
	fmt.Println("Host: ", res.Host)
	fmt.Println("Puerto; ", res.Port)
	fmt.Println("Número random: ", res.RandNumber)

	conn.Close()

	// const (
	// 	HOST = res.Host
	// 	PORT = res.Port
	// 	TYPE = "tcp"
	// )

	fmt.Println("Starting Server!")
	tcpServer, err := net.ResolveTCPAddr("tcp", res.Host+res.Port)
	if err != nil {
		fmt.Println("Unable to resolve the given address\nERROR: " + err.Error())
		os.Exit(1)
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		fmt.Printf("Unable to make a connection to the Server.\nERROR: %v", err.Error())
		os.Exit(1)
	}

	//fmt.Println(string(res.RandNumber))

	numStr := strconv.Itoa(res.RandNumber)
	fmt.Println(numStr)
	_, err = tcpConn.Write([]byte(numStr))
	if err != nil {
		fmt.Printf("Could not write the Message.\nERROR: %v", err.Error())
		os.Exit(1)
	}

	recivedBuffer := make([]byte, 1024)
	n, err = tcpConn.Read(recivedBuffer)
	fmt.Println("HOLA")
	if err != nil {
		fmt.Printf("Could not recieve data sent from the server.\nERROR: %v", err.Error())
		os.Exit(1)
	}

	receivedMessage := string(recivedBuffer[:n])
	fmt.Printf("Recieved Message Content: %v", receivedMessage)

	tcpConn.Close()
}
