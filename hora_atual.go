package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// aceita uma conexão criada por um cliente
		conn, err := listener.Accept()
		if err != nil {
			// falhas na conexão. p.ex abortamento
			log.Print(err)
			continue
		}
		// serva a conexão estabelecida
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// envia o conteúdo servido na conexão
		_, err := io.WriteString(c, time.Now().Format("02:05:00\n"))
		if err != nil {
			// p.ex erro ao enviar os dados para um cliente que desconectou
			return
		}
		time.Sleep(1 * time.Second)
	}
}
