package messenger

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/go-nats"
)

func NewClockPublisher(urls, subj string) error {

	nc, err := nats.Connect(urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for {
		msg := time.Now().Format(time.RFC3339 + "\n")
		nc.Publish(subj, []byte(msg))
		log.Printf("Published [%s] : '%s'\n", subj, msg)
		nc.Flush()
		time.Sleep(1 * time.Second)
		if err := nc.LastError(); err != nil {
			return err
		}
	}

}

func NewPublisher(urls string, subj, msg string) {

	nc, err := nats.Connect(urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Publish(subj, []byte(msg))
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[Nats clock pub] ")
}
