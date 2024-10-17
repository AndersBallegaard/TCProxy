package main

import (
	"net"
	"strconv"

	"gopkg.in/yaml.v3"
)

type ConfigServer struct {
	Port uint16
}
type ConfigTarget struct {
	Address string
	Port    uint16
}

type Config struct {
	Server ConfigServer
	Target ConfigTarget
}

func loadConfig(path string) Config {
	configFile := readFileByte(path)
	var config Config
	yaml.Unmarshal(configFile, &config)
	return config
}

func (t *ConfigTarget) getTCPAddr() *net.TCPAddr {
	addr := t.Address + ":" + strconv.Itoa(int(t.Port))
	TCPAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkErrFatal(err)
	return TCPAddr
}
