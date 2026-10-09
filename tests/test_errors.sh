
#!/bin/sh

set -eu

echo "Test 2: error handling"

output=$(printf 'cd /tmp /home\nunknown-command\nexit\n' | ../run.sh 2>&1)

echo "$output"

echo "$output" | grep -Fq 'cd: too many arguments' || {
    echo "FAIL: cd argument error was not reported"
    exit 1
}

echo "$output" | grep -Fq 'unknown-command: command not found' || {
    echo "FAIL: unknown command error was not reported"
    exit 1
}

echo "PASS: error handling"
