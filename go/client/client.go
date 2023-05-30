package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net"
	"os"
)

const (
	clientCertFile = "../../certs/client.crt"
	clientKeyFile  = "../../certs/client.key"
	serverCACert   = "../../certs/server.crt"
)

func main() {
	// Load client certificates
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatal("Error loading client certificates:", err)
	}

	// Load server CA certificate
	caCert, err := os.ReadFile(serverCACert)
	if err != nil {
		log.Fatal("Error reading server CA certificate:", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	// Connect to the server
	conn, err := tls.Dial("tcp", "localhost:12345", tlsConfig)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	log.Println("Connected to server")

	// Handle server connection
	handleServer(conn)

}

func handleServer(conn net.Conn) {
	defer conn.Close()

	log.Println("Handling server connection")

	// Write data to server
	n, err := conn.Write([]byte("Hello from client\n"))
	if err != nil {
		log.Println("Error writing to server:", err)
		return
	}

	log.Printf("Wrote %d bytes to server\n", n)

	// Read data from server
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println("Error reading from server:", err)
		}
		return
	}

	log.Printf("Read %d bytes from server: %s\n", n, buf[:n])
}
