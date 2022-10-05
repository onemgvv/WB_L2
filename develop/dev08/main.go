package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/onemgvv/WB_L2/develop/dev08/pkg"
)

type Command string

const (
	Help Command  = "/help"
	Quit 					= "/quit"
	QuitSecond    = "/q"
	CD						= "cd"
	PWD  					= "pwd"
	Echo					= "echo"
  Kill					= "kill"
  PS						= "ps"
)

func QuitFn(interface{}) {
	fmt.Println("Exit.")
	os.Exit(0)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Golang Shell\n\nTry /help to see commands\n\nEnter /quit or /q for exit\n")

	for {
		var cmd []string

		fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

		cmd = strings.Split(text, " ")

		switch cmd[0] {
		default:
			fmt.Printf("command %s not found\n", text)
		case string(Help):
			pkg.HelpCommand()
		case string(Quit):
		case string(QuitSecond):
			fmt.Println("Exit")
			os.Exit(0)
		case string(CD):
			pkg.CdCommand(cmd[1:])
		case string(PWD):
			pkg.PwdCommand()
		case string(PS):
			pkg.PsCommand(cmd[1:])
		case string(Echo):
			pkg.EchoCommand(cmd[1:])
		case string(Kill):
			pkg.KillCommand(cmd[1:])
		}
	}
}
