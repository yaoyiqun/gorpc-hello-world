package server

import (
	"google.golang.org/grpc"
	"grpc-hello-world/pkg/util"
	pb "grpc-hello-world/proto"
	"log"
	"net"
)

var (
	ServerPort     string
	CertServerName string
	CertPemPath    string
	CertKeyPath    string
	EndPoint       string
	CAPath         string
)

func Run() (err error) {
	EndPoint = ":" + ServerPort

	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP listen err:%v\n", err)
	}

	srv := grpc.NewServer(grpc.Creds(util.GetServerCreds(CertPemPath, CertKeyPath, CAPath)))
	pb.RegisterHelloWorldServer(srv, NewHelloService())
	log.Printf("gRPC and https listen on: %s\n", ServerPort)

	if err = srv.Serve(conn); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}
	return nil
}
