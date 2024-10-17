package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func (instance *Config) handleConnection(clientConn *net.TCPConn) {
	defer clientConn.Close()
	targetTCPAddr := instance.Target.getTCPAddr()
	log.Printf("Serving %s to %s\n", clientConn.RemoteAddr().String(), targetTCPAddr)

	targetConn, err := net.DialTCP("tcp", nil, targetTCPAddr)
	checkErrFatal(err)
	defer targetConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(clientConn, targetConn)
		clientConn.CloseWrite()
	}()
	go func() {
		defer wg.Done()
		io.Copy(targetConn, clientConn)
		clientConn.CloseWrite()
	}()

	wg.Wait()
}

func (instance *Config) server() {
	target := instance.Server.getTCPAddr()
	l, err := net.ListenTCP("tcp", target)
	log.Println("Listing on", target, "Forwarding connections to", instance.Target.getTCPAddr())
	checkErrFatal(err)
	defer l.Close()
	for {
		c, err := l.AcceptTCP()
		checkErrFatal(err)

		go instance.handleConnection(c)
	}
}
