package webhook

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Webhook struct {
	ID          uint
	Name        string
	Description string
	Executor    string
	Url         string
	Secret      string
}

func (w Webhook) Execute() {
	go func() {
		f, err := os.OpenFile("./log/"+strconv.Itoa(int(w.ID))+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
		cmd := exec.Command(w.Executor)
		stdout, err := cmd.Output()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		log.Println(string(stdout))
	}()
}

func (w Webhook) Print() {
	fmt.Printf("%+v\n", w)
}

func init() {
	logPath := "./log"
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		_ = os.Mkdir(logPath, os.ModeDir)
	}
}
