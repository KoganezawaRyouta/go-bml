package batche

import (
	"os"

	"github.com/KoganezawaRyouta/go-bml/config"

	kitLog "github.com/go-kit/kit/log"
)

type Person struct {
	mouth  string
	logger kitLog.Logger
}

// NewHello  init of Person
func NewHello(config *config.Config, speak string) *Person {
	person := Person{}
	person.mouth = speak

	//logfile, err := os.OpenFile("./tmp/batch.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	panic("cannot open log file:"  + config.LogFile)
	//}
	//defer logfile.Close()

	var logger kitLog.Logger
	logger = kitLog.NewJSONLogger(kitLog.NewSyncWriter(os.Stderr))
	logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC, "caller", kitLog.DefaultCaller)
	person.logger = logger

	return &person
}

// Vacuum it obtains the information of the trades and ticker from coincheck.jp,
// and register to DB
func (p *Person) Run() {
	p.logger.Log("process", "start")
	p.logger.Log("mouth", p.mouth)
	p.logger.Log("process", "end")
}
