package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func makePrompt(custom string) string {
	currentUser, err := user.Current()
	username := ""

	if err == nil {
		username = currentUser.Username
	}

	hostname, _ := os.Hostname()
	cwd, _ := os.Getwd()

	if custom == "" {
		return fmt.Sprintf("(EMU)%s@%s:%s$ ", username, hostname, cwd)
	}

	return strings.NewReplacer(
		"{user}", username,
		"{host}", hostname,
		"{cwd}", cwd,
	).Replace(custom)
}
