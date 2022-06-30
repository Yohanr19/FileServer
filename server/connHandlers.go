package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handleSubscriber(conn net.Conn) {
	ch := scanLine(conn)
	ConnMap.Add(ch, conn)
	fmt.Printf("New Subscriber on channel %s \n", ch)
}

func handlePoster(conn net.Conn) {
	defer conn.Close()

	ch := scanLine(conn)
	
	defer func() {
		for _, c := range ConnMap.Get(ch) {
			c.Close()
		}
		ConnMap.Set(ch, nil)
	}()
	conns := ConnMap.Get(ch)
	//TODO Remove all connections that are closed at this time
	ws := make([]io.Writer, len(conns))
	for i, c := range conns {
		ws[i] = c
	}
	writers := io.MultiWriter(ws...)

	_, err := io.Copy(writers, conn)
	//fmt.Println(z)
	if err != nil {
		log.Printf("Error sending in channel %s \n", ch)
		return
	}

}

//scanLine takes a reader, scans the first line and returns the text of the line
func scanLine(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text()
}
