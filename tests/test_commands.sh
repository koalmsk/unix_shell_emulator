
#!/bin/sh

set -eu

echo "Test 1: REPL commands"

output=$(printf 'ls\nls $HOME\ncd /tmp\nexit\nunknown-after-exit\n' | ../run.sh 2>&1)

echo "$output"

echo "$output" | grep -Fq 'ls: args=[]' || {
    echo "FAIL: ls without arguments"
    exit 1
}

echo "$output" | grep -Fq "ls: args=[\"$HOME\"]" || {
    echo "FAIL: environment variable HOME was not expanded"
    exit 1
}

echo "$output" | grep -Fq 'cd: args=["/tmp"]' || {
    echo "FAIL: cd stub"
    exit 1
}

if echo "$output" | grep -Fq 'unknown-after-exit'; then
    echo "FAIL: commands continued after exit"
    exit 1
fi

echo "PASS: REPL commands"
