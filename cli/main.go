package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
	"webhook-service/webhook"
)

var reader *bufio.Reader

func getUserInput(prompt string) (s string) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSuffix(s, "\r")
	return
}

func getUUID() (u string) {
	u = uuid.New().String()
	u = strings.Replace(u, "-", "", -1)
	return
}

func prompt() string {
	return "--> "
}

func shell() {
	fmt.Println(`Welcome to the cli of webhook service, type "h" to get help.`)
	for {
		fmt.Print(prompt())
		commandString := getUserInput("")
		command := strings.Split(commandString, " ")
		if executeCommand(command) {
			os.Exit(0)
		}
	}
}

func executeCommand(command []string) (stop bool) {
	stop = false
	parameter := ""
	var id uint = 0
	if len(command) > 1 {
		parameter = command[1]
		n, _ := strconv.Atoi(parameter)
		id = uint(n)
	}
	switch command[0] {
	case "":
		return
	case "n":
		fallthrough
	case "new":
		var w webhook.Webhook
		w.ID = id
		w.Url = getUUID()
		updateWebhookFromUser(&w)
		webhook.New(&w)
	case "m":
		fallthrough
	case "modify":
		w, found := webhook.Read(id)
		if found {
			updateWebhookFromUser(&w)
			webhook.Update(&w)
		} else {
			fmt.Println("Record not found.")
		}
	case "e":
		fallthrough
	case "execute":
		w, found := webhook.Read(id)
		if found {
			w.Execute()
		} else {
			fmt.Println("Record not found.")
		}
	case "l":
		fallthrough
	case "list":
		webhooks := webhook.All()
		for _, w := range webhooks {
			w.Print()
		}
	case "s":
		fallthrough
	case "search":
		webhooks := webhook.Search(parameter)
		for _, w := range webhooks {
			w.Print()
		}
	case "d":
		fallthrough
	case "delete":
		webhook.Delete(id)
	case "p":
		fallthrough
	case "print":
		w, found := webhook.Read(id)
		if found {
			w.Print()
		} else {
			fmt.Println("Record not found.")
		}
	case "q":
		fallthrough
	case "quit":
		stop = true
		return
	default:
		fmt.Println(`Help information:
1. n / new: create new webhook.
2. d / delete id: delete specified webhook.
3. m / modify id: modify existed webhook.
4. e / execute id: execute specified webhook.
5. l / list: list all webhooks.
6. s / search keyword: search webhooks by a keyword in name or description.
7. p / print id: print detail information of specified webhooks.
8. h / help: print help information. 
9. q / quit: quit cli shell.`)
	}
	return
}

func updateWebhookFromUser(w *webhook.Webhook) {
	temp := getUserInput(fmt.Sprintf("Name (current value is %q): ", w.Name))
	if temp != "" {
		w.Name = temp
	}
	temp = getUserInput(fmt.Sprintf("Description (current value is %q): ", w.Description))
	if temp != "" {
		w.Description = temp
	}
	temp = getUserInput(fmt.Sprintf("Executor Path (current value is %q): ", w.Executor))
	if temp != "" {
		w.Executor = temp
	}
	temp = getUserInput(fmt.Sprintf("Secret (current value is %q): ", w.Secret))
	if temp != "" {
		w.Secret = temp
	}
}

func main() {
	if len(os.Args) == 1 {
		shell()
	} else {
		executeCommand(os.Args[1:])
	}
}
