package clock

import (
	"io"
	"net"
	"time"

	"log"
)

func Run() error {
	listener, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		log.Fatal(err)
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return err
		}
		handleCon(conn)
	}
}

func handleCon(c net.Conn) error {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.RFC3339+"\n"))
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}
