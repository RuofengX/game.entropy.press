package jormungandr_test

import (
	service "jormungandr/v2"
	"log"
	"net"
	"testing"
)

const (
	Address string = "127.0.0.1:8089"
)

func TestServe(t *testing.T) {
	// 开启监听
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("!! Failed to listen: %v", err)
	}

	// 实例化服务
	s := service.NewService()
	log.Println("~~ 服务器启动")

	s.Start(listen, true)
}
