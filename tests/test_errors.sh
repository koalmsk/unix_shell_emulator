#!/bin/sh

echo "Test 2: script stops on error"

output=$(../run.sh --script-path ../scripts/error.sh 2>&1)

echo "$output"

echo "$output" | grep -q "unknown-command: command not found" || {
    echo "FAIL: unknown command error was not found"
    exit 1
}

echo "$output" | grep -q "script stopped at line" || {
    echo "FAIL: script did not stop after error"
    exit 1
}

if echo "$output" | grep -q "this-command-should-not-run"; then
    echo "FAIL: command after error was executed"
    exit 1
fi

echo "PASS"