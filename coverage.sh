#!/bin/sh
set -eu
ARGS=${@:-'./...'}
mkdir -p ./_testing
./test.sh -coverprofile ./_testing/coverage.out ${ARGS}
exec go tool cover -html ./_testing/coverage.out -o ./_testing/coverage.html
