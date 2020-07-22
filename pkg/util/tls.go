package util

import (
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/http2"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

//获取tls配置
func GetTLCConfig(certPemPath, certKeyPath string) *tls.Config {
	var certKeyPair *tls.Certificate
	cert, _ := ioutil.ReadFile(certPemPath)
	key, _ := ioutil.ReadFile(certKeyPath)

	//从一堆pem编码的数据中解析公钥/私钥对.成功则返回公钥私钥对
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("TLS keyPair err:%v\n", err)
	}
	certKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		//NextProtoTLS是谈判期间的NPN/ALPN协议,用于HTTP/2的TLS设置
		NextProtos: []string{http2.NextProtoTLS},
	}
}

//获取tls配置
func GetTLCConfigFromCA(certPemPath, certKeyPath, caPath string) *tls.Config {
	var certKeyPair *tls.Certificate
	cert, _ := ioutil.ReadFile(certPemPath)
	key, _ := ioutil.ReadFile(certKeyPath)
	//从一堆pem编码的数据中解析公钥/私钥对.成功则返回公钥私钥对
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("TLS keyPair err:%v\n", err)
	}
	certKeyPair = &pair

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		log.Fatalf("ioutil.ReadFile err :%v\n", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM(ca) err")
	}
	return &tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		//NextProtoTLS是谈判期间的NPN/ALPN协议,用于HTTP/2的TLS设置
		//NextProtos: []string{http2.NextProtoTLS},
	}
}

func NewTLSListener(inner net.Listener, config *tls.Config) net.Listener {
	return tls.NewListener(inner, config)
}

//返回服务端的creds
func GetServerCreds(certPemPath, certKeyPath, caPath string) credentials.TransportCredentials {
	var certKeyPair *tls.Certificate
	cert, _ := ioutil.ReadFile(certPemPath)
	key, _ := ioutil.ReadFile(certKeyPath)
	//从一堆pem编码的数据中解析公钥/私钥对.成功则返回公钥私钥对
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("TLS keyPair err:%v\n", err)
	}
	certKeyPair = &pair

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		log.Fatalf("ioutil.ReadFile err :%v\n", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM(ca) err")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	return creds
}

//返回客户端的creds
func GetClientCreds(certPemPath, certKeyPath, caPath, certServerName string) credentials.TransportCredentials {
	var certKeyPair *tls.Certificate
	cert, _ := ioutil.ReadFile(certPemPath)
	key, _ := ioutil.ReadFile(certKeyPath)
	//从一堆pem编码的数据中解析公钥/私钥对.成功则返回公钥私钥对
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("TLS keyPair err:%v\n", err)
	}
	certKeyPair = &pair

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		log.Fatalf("ioutil.ReadFile err :%v\n", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM(ca) err")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		ServerName:   certServerName,
		RootCAs:      certPool,
	})

	return creds
}
