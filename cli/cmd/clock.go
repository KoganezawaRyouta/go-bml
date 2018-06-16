package cmd

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/go-bml/clock"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		log.SetOutput(os.Stdout)
		log.SetPrefix("[Clock] ")
		go func() {
			pid := os.Getpid()
			pidInfo, _ := ps.FindProcess(pid)
			log.Printf("start server!!")
			log.Printf(" PID          : %d\n", pidInfo.Pid())
			log.Printf(" PPID         : %d\n", pidInfo.PPid())
			log.Printf(" Process name : %s\n", pidInfo.Executable())
			pp, _ := ps.FindProcess(pidInfo.PPid())
			log.Printf(" Parent process name : %s\n", pp.Executable())
			errsCh <- clock.Run()
		}()
		log.Fatal("terminated", <-errsCh)
	},
}

func init() {
	RootCmd.AddCommand(clockCmd)
}
