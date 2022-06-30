package main

import (
	"net"
	"sync"
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

//Takes a ch and removes all connections that have been closed
func (sm *SafeMap) CloseConnections() {
	for _, conns := range sm.connMap {
		for _, c := range conns {
			c.Close()
		}
	}
}
