package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleSubscriber(conn net.Conn, ch string) error {
	defer conn.Close()
	var chBuf = make([]byte, CHANNEL_BYTES)
	var filnameBuf = make([]byte, FILENAME_BYTES)
	copy(chBuf, ch)
	conn.Write(chBuf)
	conn.Read(filnameBuf)
	filename := string(trim(filnameBuf))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	var totalR int
	var totalW int
	for {
		var buf = make([]byte, 1<<20)
		n, err := conn.Read(buf)
		totalR += n
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		wrote, _ := f.Write(buf[:n])
		totalW += wrote
	}
	/* n, err := io.Copy(f, conn)
	if err != nil {
		return err
	} */
	fmt.Printf("Transfered %q, Read %d bytes, wrote %d bytes \n", filename, totalR, totalW)
	return nil
}
func handlePoster(conn net.Conn, filename string, ch string) error {
	defer conn.Close()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	var chBuf = make([]byte, CHANNEL_BYTES)
	var filnameBuf = make([]byte, FILENAME_BYTES)
	copy(chBuf, ch)
	copy(filnameBuf, filename)
	conn.Write(chBuf)
	conn.Write(filnameBuf)
	fmt.Println("Wrote headers")
	n, err := io.Copy(conn, f)
	if err != nil {
		return err
	}
	fmt.Printf("Sent file %q with size of %d bytes \n", filename, n)
	return nil
}

//takes a slyce of bytes and returns the slice with all trailing 0 elements trimed down
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
