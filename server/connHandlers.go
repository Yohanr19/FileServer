package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//Make nil when closing conn, before adding to multiwriter check if is nil and dont add it if is nil
func handleSubscriber(conn net.Conn) {
	var chBuf = make([]byte, CHANNEL_BYTES)
	conn.Read(chBuf)
	ch := string(chBuf)
	ConnMap.Add(ch, conn)
	fmt.Printf("New Subscriber on channel %s \n", ch)
}

func handlePoster(conn net.Conn) {
	defer conn.Close()

	var chBuf = make([]byte, CHANNEL_BYTES)
	conn.Read(chBuf)
	ch := string(chBuf)

	defer func() {
		for _, c := range ConnMap.Get(ch) {
			if c != nil {
				c.Close()
			}
		}
		ConnMap.Set(ch, nil)
	}()
	conns := ConnMap.Get(ch)
	//TODO Remove all connections that are closed at this time
	for i, c := range conns {
		var one = make([]byte, 1)
		err := c.SetReadDeadline(time.Now().Add(time.Millisecond * 15))
		if err != nil {
			log.Print(err)
		}
		_, err = c.Read(one)
		if err == io.EOF {
			c.Close()
			conns[i] = nil
		}
		log.Print(err)
	}
	ws := make([]io.Writer, 0)
	for _, c := range conns {
		//Does not add nil connections
		if c != nil {
			ws = append(ws, c)
		}
	}
	fmt.Println(len(ws))
	writers := io.MultiWriter(ws...)

	z, err := io.Copy(writers, conn)
	fmt.Println(z)
	if err != nil {
		log.Printf("Error sending in channel %s \n", ch)
		return
	}
}
