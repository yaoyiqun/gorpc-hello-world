package util

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
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
