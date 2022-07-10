package main

import (
	"net"
	"sync"
	"time"
	"io"
	"log"
)

type SafeMap struct {
	connMap map[string][]net.Conn
	mu      sync.Mutex
}

func NewSafeMap() *SafeMap {
	sm := SafeMap{}
	sm.mu = sync.Mutex{}
	sm.connMap = make(map[string][]net.Conn)
	return &sm
}

func (sm *SafeMap) Add(ch string, conn net.Conn) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.connMap[ch] = append(sm.connMap[ch], conn)
}

func (sm *SafeMap) Get(ch string) []net.Conn {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.connMap[ch]
}

func (sm *SafeMap) Set(ch string, conns []net.Conn) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.connMap[ch] = conns
}
func (sm *SafeMap) CloseConnections() {
	for _, conns := range sm.connMap {
		for _, c := range conns {
			if c != nil {
				c.Close()
			}
		}
	}
}
func (sm *SafeMap)ClearConnections(){
	for _, conns := range sm.connMap{
		if conns!=nil{
			clearConns(conns)
		}
	}
}
//Returns a copy of the current Map of connections
func (sm *SafeMap)CopyMap()map[string][]net.Conn{
	res := make(map[string][]net.Conn) 
	for ch , conn := range sm.connMap {
		if len(conn) > 0{
		sm.mu.Lock()
		res[ch] = conn
		sm.mu.Unlock()	
		}
	}
	return res
}
//takes an array of connections, and will remove all closed connections
func clearConns(conns []net.Conn){
	for i, c := range conns {
		if c == nil {
			continue
		}
		var one = make([]byte, 1)
		err := c.SetReadDeadline(time.Now().Add(time.Millisecond * 15))
		if err != nil {
			log.Print(err)
			continue
		}
		_, err = c.Read(one)
		if err == io.EOF {
			c.Close()
			conns[i] = nil
		}
	}
}
