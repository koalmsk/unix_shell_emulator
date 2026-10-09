package main

import (
	"fmt"
	"os"
	"os/user"
)

func makePrompt() string {
	currentUser, err := user.Current()
	username := ""

	if err == nil {
		username = currentUser.Username
	}

	hostname, _ := os.Hostname()
	cwd, _ := os.Getwd()

	return fmt.Sprintf("(EMU)%s@%s:%s$ ", username, hostname, cwd)

}
