package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

func (instance *Config) handleConnection(clientConn net.Conn) {
	//defer clientConn.Close()
	targetTCPAddr := instance.Target.getTCPAddr()
	log.Printf("Serving %s\n", clientConn.RemoteAddr().String(), "to", targetTCPAddr)

	targetConn, err := net.DialTCP("tcp", nil, targetTCPAddr)
	//defer targetConn.Close()
	checkErrFatal(err)
	for {
		cReader := bufio.NewReader(clientConn)
		cBytes := make([]byte, 0xffff)
		_, err := cReader.Read(cBytes)
		checkErrFatal(err)

		_, err = targetConn.Write(cBytes)
		checkErrFatal(err)

		tReader := bufio.NewReader(targetConn)
		tBytes := make([]byte, 1024)
		_, err = tReader.Read(tBytes)
		checkErrFatal(err)

		clientConn.Write(tBytes)
	}
}

func (instance *Config) server() {
	l, err := net.Listen("tcp", "[::]:"+strconv.Itoa(int(instance.Server.Port)))
	checkErrFatal(err)
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go instance.handleConnection(c)
	}
}
