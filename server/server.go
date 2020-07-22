package server

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-hello-world/pkg/util"
	pb "grpc-hello-world/proto"
	"log"
	"net"
	"net/http"
)

var (
	ServerPort  string
	CertName    string
	CertPemPath string
	CertKeyPath string
	EndPoint    string
)

func Server() (err error) {
	EndPoint = ":" + ServerPort
	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP listen err:%v\n", err)
	}

	tlsConfig := util.GetTLCConfig(CertPemPath, CertKeyPath)
	srv := createInternalServer(conn, tlsConfig)

	log.Printf("gRPC and https listen on: %s\n", ServerPort)

	if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}
	return nil
}

func createInternalServer(conn net.Listener, tlsconfig *tls.Config) *http.Server {

	var opts []grpc.ServerOption

	//构造TLS证书凭证
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v\n", err)
	}
	//服务器连接设置凭据
	opts = append(opts, grpc.Creds(creds))
	//创建gRPC服务器
	grpcServer := grpc.NewServer(opts...)
	//注册gRPC服务
	pb.RegisterHelloWorldServer(grpcServer, NewHelloService())

	//创建grpc-gateway关联组件
	ctx := context.Background()
	//从客户机的输入证书文件构造TLS凭证
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v\n", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	//注册具体的服务
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsconfig,
	}
}
