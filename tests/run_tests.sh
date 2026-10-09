#!/bin/sh

CDPATH="" cd -- "$(dirname -- "$0")" || exit 1

echo "Running tests..."
echo

./test_commands.sh || exit 1
echo

./test_errors.sh || exit 1
echo

./test_startup.sh || exit 1
echo

echo "All tests passed!"