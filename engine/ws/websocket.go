package main

import (
	"github.com/gobwas/ws"
	"io"
	"log"
	"net"
	"time"
)

type Connect struct {
	net.Conn
	t time.Duration
}

func main() {
	log.Println("Scyna WS Start with port 8080")

	timeout := time.Duration(100 * time.Millisecond)
	address := "localhost:8080"

	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Can not create listener " + err.Error())
		return
	}

	for {
		conn, errListener := ln.Accept()
		if errListener != nil {
			log.Println("Can not accept connection ws " + errListener.Error())
		}
		log.Println("Connection accept - " + ln.Addr().String())

		_, errUpgrade := ws.Upgrade(&Connect{conn, timeout})
		if errUpgrade != nil {
			log.Println("Can not upgrade connection ws " + errUpgrade.Error())
		}

		go func() {
			defer conn.Close()

			for {
				header, err := ws.ReadHeader(conn)
				if err != nil {
					// handle error
				}

				payload := make([]byte, header.Length)
				_, err = io.ReadFull(conn, payload)
				if err != nil {
					// handle error
				}
				if header.Masked {
					ws.Cipher(payload, header.Mask, 0)
				}

				// Reset the Masked flag, server frames must not be masked as
				// RFC6455 says.
				header.Masked = false

				if err := ws.WriteHeader(conn, header); err != nil {
					// handle error
				}
				if _, err := conn.Write(payload); err != nil {
					// handle error
				}

				if header.OpCode == ws.OpClose {
					log.Println("WS Close")
					return
				}
			}

		}()
	}

}
