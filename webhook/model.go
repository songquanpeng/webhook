package webhook

import (
	"fmt"
	"os/exec"
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
	go exec.Command(w.Executor)
}

func (w Webhook) Print() {
	fmt.Printf("%+v\n", w)
}
