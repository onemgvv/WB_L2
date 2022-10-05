package pkg

import (
	"fmt"
	"os/exec"
	"strings"
)

func HelpCommand()  {
	fmt.Println()
	fmt.Println("============= [HELP] =============")
	fmt.Println()
	fmt.Println("1. /help       - see all commands")
	fmt.Println("2. /quit(/q)   - exit")
	fmt.Println("3. cd <args>   - change directory")
	fmt.Println("4. pwd         - print work directory")
	fmt.Println("5. echo <args> - print something in stdout")
	fmt.Println("6. kill <args> - kill process")
	fmt.Println("7. ps <args>   - info about active processes")
	fmt.Println()
}

func CdCommand(args []string) {
	byte, err := exec.Command("cd", args...).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(strings.Replace(string(byte), "\n", "", -1))
}

func PwdCommand() {
	byte, err := exec.Command("pwd").CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(strings.Replace(string(byte), "\n", "", -1))
}

func PsCommand(args []string) {
	byte, err := exec.Command("ps", args...).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(strings.Replace(string(byte), "\n", "", -1))
}

func EchoCommand(args []string) {
	byte, err := exec.Command("echo", args...).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(strings.Replace(string(byte), "\n", "", -1))
}

func KillCommand(args []string) {
	byte, err := exec.Command("kill", args...).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(strings.Replace(string(byte), "\n", "", -1))
}