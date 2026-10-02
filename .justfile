default: build serve

build:
  GOOS=js GOARCH=wasm go build -o go.wasm .

serve:
  #! /bin/sh

  if command -v python3 >/dev/null 2>&1; then
    python -m http.server -d .
  else
    echo 'Python is required to run this dev server'
    exit 1
  fi

ci:
  ./.ci.sh
