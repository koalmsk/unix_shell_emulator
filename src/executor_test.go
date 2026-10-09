package main

import (
	"testing"
	
)

func TestExecuteLS(t *testing.T) {
	result := execute("ls")

	if result != ok {
		t.Fatalf("expected ok, got %d", result)
	}
}

func TestExecuteUnknownCommand(t *testing.T) {
	result := execute("unknown-command")

	if result != commandError {
		t.Fatalf("expected commandError, got %d", result)
	}
}

func TestExecuteExit(t *testing.T) {
	result := execute("exit")

	if result != exit {
		t.Fatalf("expected exit, got %d", result)
	}
}

func TestExecuteCdTooManyArguments(t *testing.T) {
	result := execute("cd /tmp /home")

	if result != commandError {
		t.Fatalf("expected commandError, got %d", result)
	}
}
