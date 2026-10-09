package main

import (
	"bufio"
	"fmt"
	"os"
)

func runInteractive(prompt string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			break
		}

		result := execute(scanner.Text())

		if result == exit {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("input error:", err)
	}
}

func main() {

	prompt := makePrompt()

	runInteractive(prompt)
}
