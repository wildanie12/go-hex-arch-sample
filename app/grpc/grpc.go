package grpc

import (
	"fmt"
	"log"
	"net"

	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
	"google.golang.org/grpc"
)

// MainGRPC struct
type MainGRPC struct {
	host string
	port string
	server *grpc.Server
}

// New main gRPC Instance
func New(host, port string) *MainGRPC {
	return &MainGRPC{
		host: host,
		port: port,
	}
}

// Start main gRPC server
func (rpc *MainGRPC) Start() {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", rpc.host, rpc.port))
	if err != nil {
		log.Println(_color.ThisF("red", "[gRPC] failed to listen: %v", err))
		return
	}

	provider := rpc.provide()

	rpc.server = grpc.NewServer()
	rpc.defineRoute(provider)
	
	log.Println(_color.This("green", "|-----------------------------------------------"))
	log.Println(_color.ThisF("green", "| gRPC listening on %s:%s", rpc.host, rpc.port))
	log.Println(_color.This("green", "|-----------------------------------------------"))
	
	err = rpc.server.Serve(listener)
	if err != nil {
		log.Println(_color.ThisF("red", "[gRPC] failed to serve grpc on %s:%s : %v", rpc.host, rpc.port, err))
		return
	}
}