package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/nats-io/nats.go"
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
	if len(args) != 5 {
		fmt.Println("Please provide listen host & port and NATS host and subject")
		fmt.Println("\tExample: localhost 1234 nats://demo.nats.io:4222 my-subject")
		return
	}
	fmt.Println("Listening to ", LISTEN_HOST, ":", LISTEN_PORT)
	fmt.Println("Publishing to ", NATS_HOST, " on the subject ", NATS_SUBJECT)

	conn, err := net.Dial(LISTEN_TYPE, LISTEN_HOST+":"+LISTEN_PORT)
	if err != nil {
		fmt.Println("There was an error connecting =(")
		fmt.Println(err)
	}
	nc, err := nats.Connect(NATS_HOST)
	if err != nil {
		fmt.Println("Error connecting to nats:")
		fmt.Println(err)
	}

	fmt.Println("Connections successful, listening now...")

	scanner := bufio.NewScanner(conn)
	fmt.Print("\033[s") // Set the cursor position
	var i int = 0
	for scanner.Scan() {
		i += 1
		fmt.Print("\033[u\033[K")
		fmt.Println("Messages Handled: ", i)
		// fmt.Println(scanner.Text())
		// fmt.Println("")
		go handleMessage(nc, NATS_SUBJECT, scanner.Text())
	}
}

func handleMessage(nc *nats.Conn, subject string, msg string) {
	nc.Publish(subject, []byte(msg))
}
