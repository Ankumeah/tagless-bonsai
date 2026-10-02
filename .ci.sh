#! /bin/sh

set -eu

export GOOS=js
export GOARCH=wasm

run() {
  echo "-> ${@}"

  if output=$("${@}" 2>&1); then
    return 0
  else
    status=${?}
    echo "${output}"
    return "${status}"
  fi
}

run go vet -tags=js,wasm "${@}" ./...
run go fix -diff -tags=js,wasm "${@}" ./...

if test -n "${CI:-}"; then
  run gofmt -d .
else
  run go fmt ./...
fi
