package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runScript(filename string, prompt string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("cannot open script:", err)
		return commandError
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++

		result := runScriptLine(scanner.Text(), prompt, lineNumber)

		if result != ok {
			return result
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading script:", err)
		return commandError
	}

	return ok
}

func runScriptLine(line string, prompt string, lineNumber int) int {
	command := strings.TrimSpace(line)

	if command == "" || strings.HasPrefix(command, "#") {
		return ok
	}

	fmt.Print(prompt)
	fmt.Println(command)

	result := execute(command)

	if result == commandError {
		fmt.Printf("script stopped at line %d\n", lineNumber)
	}

	return result
}
