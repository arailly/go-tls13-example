package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"os"
)

func main() {
	caCert, err := os.ReadFile("config/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)
	w, err := os.OpenFile("/tmp/tls-secrets.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := tls.Dial("tcp", "localhost:2000", &tls.Config{
		RootCAs:      rootCAs,
		KeyLogWriter: w,
	})
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
	conn.Close()
}
