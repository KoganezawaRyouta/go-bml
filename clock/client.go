package clock

import (
	"net"
	"time"

	"log"
)

func ClientRun() error {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer conn.Close()
	for {
		conn.Write([]byte(time.Now().Format(time.RFC3339 + "\n")))
		time.Sleep(1 * time.Second)
	}
}
