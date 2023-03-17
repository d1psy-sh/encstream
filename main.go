package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var flagServer bool

func main() {
	// - client server
	// - server encryption stream
	// - client decryption stream
	// - print("start")
	setUpFlags()

	fmt.Print("start...")
	if flagServer {
		fmt.Println("server")
		server()
	} else {
		fmt.Println("client")
		client()
	}
}

func setUpFlags() {
	flag.BoolVar(&flagServer, "server", false, "run as server")
	flag.Parse()
}

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go encryptStream(conn)
		decryptStream(conn)
	}
}

func client() {
	for {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		// go io.Copy(conn, os.Stdin)
		// io.Copy(os.Stdout, conn)
		go encryptStream(conn)
		decryptStream(conn)
	}
}

func encryptStream(conn net.Conn) {
	// take stdin
	// encrypt witn aes (use base 64 encoding for now cuz this is easier)
	// write to tcp conn
	streamEncoder := base64.NewEncoder(base64.StdEncoding, conn)
	// read from stdin
	io.Copy(streamEncoder, os.Stdin)
}

func decryptStream(conn net.Conn) {
	// read from tcp conn
	// decrypt witn aes (use base 64 encoding for now cuz this is easier)
	// write to stdout
	streamDecoder := base64.NewDecoder(base64.StdEncoding, conn)
	// write to stdout
	io.Copy(os.Stdout, streamDecoder)
}
