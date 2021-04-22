package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	VERSION = "0.0.1"

	LISTEN_HOST = "192.168.1.169"
	LISTEN_PORT = "30003"
	LISTEN_TYPE = "tcp"

	NATS_HOST    = "nats://demo.nats.io:4222"
	NATS_SUBJECT = "dan-adsb"
)

func main() {
	fmt.Println("Starting Go ADS-B NATS Forwarder")

	args := os.Args
	if len(args) != 3 {
		fmt.Println("Please provide host & port.")
		fmt.Println("\tExample: localhost 1234.")
		return
	}

	conn, err := net.Dial(LISTEN_TYPE, LISTEN_HOST+":"+LISTEN_PORT)
	if err != nil {
		fmt.Println("There was an error connecting =(")
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
