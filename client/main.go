package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	subscriberAddr = "localhost:2020"
	posterAddr     = "localhost:2021"
	CHANNEL_BYTES  = 3
	FILENAME_BYTES = 256
)

func main() {
	args := os.Args[1:]
	err := validateArguments(args)
	if err != nil {
		fmt.Println(err)
		//TODO print usage
		os.Exit(1)
	}
	switch args[0] {
	case "subscribe":
		var conn net.Conn
		go func() {
			input := bufio.NewScanner(os.Stdin)
			input.Split(bufio.ScanBytes)
			input.Scan()
			err := conn.Close()
			if err != nil {
				log.Print(err)
			}
			fmt.Println("Connection closed successfully")
			os.Exit(0)
		}()
		for {
			var err error
			conn, err = net.Dial("tcp", subscriberAddr)
			if err != nil {
				log.Fatal(err)
			}
			err = handleSubscriber(conn, args[1])
			if err != nil {
				log.Fatal(err)
			}
		}

	case "post":
		conn, err := net.Dial("tcp", posterAddr)
		if err != nil {
			log.Fatal(err)
		}
		err = handlePoster(conn, args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

	}
}

func validateArguments(arg []string) error {
	if len(arg) < 1 {
		return fmt.Errorf("unsupported amount of arguments")
	}
	switch arg[0] {
	case "post":
		if len(arg) != 3 {
			return fmt.Errorf("unsupported amount of arguments")
		}
		if ch, err := strconv.Atoi(arg[2]); ch < 1 || ch > 200 || err != nil {
			return fmt.Errorf("unsupported channel")
		}
		return nil
	case "subscribe":
		if len(arg) > 3 || len(arg) < 2 {
			return fmt.Errorf("unsupported amount of arguments")
		}
		if ch, err := strconv.Atoi(arg[1]); ch < 1 || ch > 200 || err != nil {
			return fmt.Errorf("unsupported channel")
		}
		return nil
	default:
		return fmt.Errorf("unsupported command")
	}
}
