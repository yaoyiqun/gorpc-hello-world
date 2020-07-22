package client

import (
	"context"
	"google.golang.org/grpc"
	"grpc-hello-world/pkg/util"
	pb "grpc-hello-world/proto"
	"grpc-hello-world/server"
	"log"
)

var (
	CertPemPath string
	CertKeyPath string
)

func Client() {
	/*cert,err := tls.LoadX509KeyPair("./certs/client.pem","./certs/client.key")
	if err!=nil{
		log.Fatalf("tls.loadx509keypair error\n")
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile("./certs/ca.pem")
	if err!=nil{
		log.Fatalf("ioutil.ReadFile error")
	}
	if ok:=certPool.AppendCertsFromPEM(ca);!ok{
		log.Fatalf("certPool.AppendCertsFromPEM(ca) error")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "grpc.abc",
		RootCAs: certPool,
	})*/
	creds := util.GetClientCreds(CertPemPath, CertKeyPath, server.CAPath, server.CertServerName)
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer: "Grpc",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
