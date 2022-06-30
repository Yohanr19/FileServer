package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func handleSubscriber(conn net.Conn, ch string) error {
	defer conn.Close()
	io.WriteString(conn, ch+"\n")
	filename := scanLine(conn)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := io.Copy(f, conn)
	if err != nil {
		return err
	}
	fmt.Printf("Transfered %q with size of %d bytes \n", filename, n)
	return nil
}
func handlePoster(conn net.Conn, filename string, ch string) error {
	defer conn.Close()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	io.WriteString(conn, ch+"\n")
	io.WriteString(conn, f.Name()+"\n")
	fmt.Println("Wrote headers")
	n, err := io.Copy(conn, f)
	if err != nil {
		return err
	}
	fmt.Printf("Sent file %q with size of %d bytes \n", filename, n)
	return nil
}

//scanLine takes a reader, scans the first line and returns the text of the line
func scanLine(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text()
}
