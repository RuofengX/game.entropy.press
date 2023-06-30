package main

import (
	service "jormungandr/v2"
	"log"
	"net"
)

const (
	Address = "0.0.0.0:8089"
)
func main() {
	// 开启监听
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 实例化服务
	s := new(service.ContinuumService)
	s.Start(listen)
}
