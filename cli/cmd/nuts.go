package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/go-bml/messenger"
	"github.com/mitchellh/go-ps"
	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
)

var (
	server  string
	subject string
	message string
)

var subCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		messenger.NewSubscriber(server, subject)
	},
}

var pubCmd = &cobra.Command{
	Use:   "publisher",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		messenger.NewPublisher(server, subject, message)
	},
}

var pubClockCmd = &cobra.Command{
	Use:   "clock_publisher",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[Clock Publisher Server] ")
		go func() {
			pid := os.Getpid()
			pidInfo, _ := ps.FindProcess(pid)
			log.Printf("start")
			log.Printf(" PID          : %d\n", pidInfo.Pid())
			log.Printf(" PPID         : %d\n", pidInfo.PPid())
			log.Printf(" Process name : %s\n", pidInfo.Executable())
			pp, _ := ps.FindProcess(pidInfo.PPid())
			log.Printf(" Parent process name : %s\n", pp.Executable())
			errsCh <- messenger.NewClockPublisher(server, subject)
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

func init() {
	subCmd.PersistentFlags().StringVar(&server, "s", nats.DefaultURL, "")
	subCmd.PersistentFlags().StringVar(&subject, "sub", "", "")
	RootCmd.AddCommand(subCmd)

	pubCmd.PersistentFlags().StringVar(&server, "s", nats.DefaultURL, "")
	pubCmd.PersistentFlags().StringVar(&subject, "sub", "", "")
	pubCmd.PersistentFlags().StringVar(&message, "m", "", "")
	RootCmd.AddCommand(pubCmd)

	pubClockCmd.PersistentFlags().StringVar(&server, "s", nats.DefaultURL, "")
	pubClockCmd.PersistentFlags().StringVar(&subject, "sub", "", "")
	RootCmd.AddCommand(pubClockCmd)
}
