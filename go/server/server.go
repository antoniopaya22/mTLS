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
	serverCertFile = "../../certs/server.crt"
	serverKeyFile  = "../../certs/server.key"
	clientCACert   = "../../certs/client.crt"
)

func main() {
	// Load server certificates
	cert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		log.Fatal("Error loading server certificates:", err)
	}

	// Load client CA certificate
	caCert, err := os.ReadFile(clientCACert)
	if err != nil {
		log.Fatal("Error reading client CA certificate:", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS (use Common Name)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ClientCAs:          caCertPool,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		InsecureSkipVerify: true,
	}

	// Create TCP listener
	listener, err := tls.Listen("tcp", "localhost:12345", tlsConfig)
	if err != nil {
		log.Fatal("Error creating listener:", err)
	}
	defer listener.Close()

	log.Println("Server started and listening on localhost:12345")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the client connection in a separate goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	log.Println("Handling client connection from", conn.RemoteAddr())

	// Read data from client
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println("Error reading from client:", err)
		}
		return
	}

	log.Printf("Read %d bytes from client: %s\n", n, buf[:n])

	// Write data to client
	n, err = conn.Write([]byte("Hello from server\n"))
	if err != nil {
		log.Println("Error writing to client:", err)
		return
	}

	log.Printf("Wrote %d bytes to client\n", n)
}
