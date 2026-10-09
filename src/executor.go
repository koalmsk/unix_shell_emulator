package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ok = iota
	commandError
	exit
	maxCdArgs   = 2
	maxExitArgs = 1
)

func execute(command string) int {

	command = os.ExpandEnv(command)

	args := strings.Fields(command)

	if len(args) == 0 {
		return ok
	}

	switch args[0] {

	case "ls":
		fmt.Printf("ls: args=%q\n", args[1:])
		return ok

	case "cd":
		if len(args) > maxCdArgs {
			fmt.Println("cd: too many arguments")
			return commandError
		}

		fmt.Printf("cd: args=%q\n", args[1:])
		return ok

	case "exit":
		if len(args) > maxExitArgs {
			fmt.Println("exit: too many arguments")
			return commandError
		}

		return exit

	default:
		fmt.Printf("%s: command not found\n", args[0])
		return commandError
	}
}
