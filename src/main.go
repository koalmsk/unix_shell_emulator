package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type config struct {
	vfsPath      string
	customPrompt string
	scriptPath   string
}

func parseFlags() config {
	vfsPath := flag.String(
		"vfs-path",
		"",
		"path to virtual file system",
	)

	customPrompt := flag.String(
		"prompt",
		"",
		"custom prompt",
	)

	scriptPath := flag.String(
		"script-path",
		"",
		"startup script",
	)

	flag.Parse()

	return config{
		vfsPath:      *vfsPath,
		customPrompt: *customPrompt,
		scriptPath:   *scriptPath,
	}
}

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
	cfg := parseFlags()

	if cfg.vfsPath != "" {
		fmt.Println("VFS path:", cfg.vfsPath)
	}

	prompt := makePrompt(cfg.customPrompt)

	if cfg.scriptPath != "" {
		result := runScript(cfg.scriptPath, prompt)

		if result == commandError || result == exit {
			return
		}
	}

	runInteractive(prompt)
}
