package main

import (
	"fmt"
	"io"
	"log"
	"net"

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

	var report controlers.ReportData
	report.SenderAdd = conn.RemoteAddr().String()
	defer func ()  {
		var rc, err = controlers.NewReportControler()
		if err!=nil{
			log.Printf("database error: %s",err)
		}
		if report.Status == "" {
			report.Status = "Failed"
		}
		rc.AddReport(report)
	}()

	var chBuf = make([]byte, CHANNEL_BYTES)
	conn.Read(chBuf)
	ch := string(trim(chBuf))
	report.Channel = ch
	//Dump all the connections, leaving the server ready to reestablish the connections
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
	clearConns(conns)
	ws := make([]io.Writer, 0)
	for _, c := range conns {
		//Does not add nil connections
		if c != nil {
			ws = append(ws, c)
		}
	}
	report.SubscriberAmount = len(ws)
	
	if len(ws) == 0 {
		return
	}
	writers := io.MultiWriter(ws...)	

	var filenameBuf = make([]byte, FILENAME_BYTES)
	conn.Read(filenameBuf)

	_, err := writers.Write(filenameBuf)
	if err != nil {
		log.Printf("Error sending in channel %s \n", ch)
		return
	}
	report.Filename = string(trim(filenameBuf))
	z, err := io.Copy(writers, conn)
	if err != nil {
		log.Printf("Error sending in channel %s \n", ch)
		return
	}
	report.Filesize = int(z)
	report.Status = `Completed`
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