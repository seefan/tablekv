package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"bufio"
	"net/http"
	"net/url"
	"io/ioutil"
)

var (
	host    string
	port    int
	helpStr = `Usage:command [options]
help:Print the help message.
show tables:show table list.
`
	httpUrl string
)

const BadCommand = "Incorrect command format"

func main() {
	host = *flag.String("host", "127.0.0.1", "host 127.0.0.1")
	port = *flag.Int("port", 8080, "port 8080")
	fmt.Println("Please enter command, such as help")
	httpUrl = fmt.Sprintf("http://%s:%d", host, port)
	inputCommand()
}
func inputCommand() {
	var cmd string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cmd = input.Text()
		if len(cmd) == 0 {
			println("Please enter command.")
			continue
		}
		cmds := strings.Split(strings.TrimSpace(cmd), " ")
		switch cmds[0] {
		case "help":
			runHelp()
		case "exit", "quit":
			os.Exit(0)
		default:
			runCmd(cmds)
		}
	}
}
func runHelp() {
	fmt.Println(helpStr)
}
func runCmd(cmd []string) {
	if len(cmd) == 0 {
		fmt.Println(BadCommand)
	} else {
		param := make(url.Values)
		param["cmd"] = cmd
		if r, err := http.PostForm(httpUrl, param); err == nil {
			defer r.Body.Close()
			if bs, err := ioutil.ReadAll(r.Body); err == nil {
				fmt.Print(string(bs))
			} else {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}
