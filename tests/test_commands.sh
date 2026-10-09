#!/bin/sh

echo "Test 1: basic commands"

output=$(printf "ls\nexit\n" | ../run.sh 2>&1)

echo "$output"

echo "$output" | grep -q "ls:" || {
    echo "FAIL: ls was not executed"
    exit 1
}

echo "PASS"