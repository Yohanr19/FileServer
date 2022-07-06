package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/yohanr19/fileserver/server/pkg/controlers"
)

func handleSubscriber(conn net.Conn) {
	var chBuf = make([]byte, CHANNEL_BYTES)
	conn.Read(chBuf)
	ch := string(trim(chBuf))
	ConnMap.Add(ch, conn)
	fmt.Printf("New Subscriber on channel %s \n", ch)
}

func handlePoster(conn net.Conn) {
	defer conn.Close()

	var chBuf = make([]byte, CHANNEL_BYTES)
	conn.Read(chBuf)
	ch := string(trim(chBuf))

	defer func() {
		for _, c := range ConnMap.Get(ch) {
			if c != nil {
				c.Close()
			}
		}
		ConnMap.Set(ch, nil)
	}()
	conns := ConnMap.Get(ch)
	// Remove all connections that are closed at this time
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
	writers := io.MultiWriter(ws...)
	var filenameBuf = make([]byte, FILENAME_BYTES)
	conn.Read(filenameBuf)

	_, err := writers.Write(filenameBuf)
	if err != nil {
		log.Printf("Error sending in channel %s \n", ch)
		return
	}

	z, err := io.Copy(writers, conn)
	fmt.Println(z)
	if err != nil {
		log.Printf("Error sending in channel %s \n", ch)
		return
	}
	var rc, _ = controlers.NewReportControler()
	var report controlers.ReportData
	report.Channel = ch
	report.Filename = string(trim(filenameBuf))
	report.Status = `Completed`
	report.Filesize = int(z)
	report.SenderAdd = conn.RemoteAddr().String()
	report.SubscriberAmount = len(ws)
	rc.AddReport(report)
}

func trim(a []byte) []byte {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != 0 {
			// found the first non-zero byte
			// return the slice from start to the index of the first non-zero byte
			return a[:i+1]
		}
	}
	return []byte{}
}
