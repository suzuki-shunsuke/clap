#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
  go test -race -coverprofile=profile.out -covermode=atomic "$d"
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done
