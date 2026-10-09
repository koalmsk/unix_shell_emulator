#!/bin/sh

echo "Test 3: startup script"

output=$(../run.sh \
    --vfs-path ../test-vfs \
    --prompt '{user}@{host}:{cwd}$ ' \
    --script-path ../scripts/start.sh 2>&1)

echo "$output"

echo "$output" | grep -q "ls:" || {
    echo "FAIL: ls was not executed"
    exit 1
}

echo "$output" | grep -q "cd:" || {
    echo "FAIL: cd was not executed"
    exit 1
}

echo "PASS"