
#!/bin/sh

set -eu

cd "$(dirname "$0")"

echo "Running REPL tests..."
echo

./test_commands.sh
echo

./test_errors.sh
echo

echo "All tests passed!"
