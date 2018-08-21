package clock

import (
	"fmt"
	"net"

	"log"
)

func ServerRun() error {
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

func handleCon(c net.Conn) {
	defer c.Close()
	fmt.Println("クライアントからの受信メッセージ:")
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
