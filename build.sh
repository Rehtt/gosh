#!/bin/bash
linux=('amd64' '386' 'arm' 'arm64' 'mipsle')
windows=('amd64' '386')
version="0.0.0"
# shellcheck disable=SC2068
for archName in ${linux[@]}; do
  export GOOS=linux
  export GOARCH=$archName
  fileName='gosh_'$version'_linux_'$archName
  go build -ldflags="-s -w" -o $fileName main.go
  upx --best --no-progress $fileName
  echo "build $archName done"
done
# shellcheck disable=SC2068
for archName in ${windows[@]}; do
  export GOOS=windows
  export GOARCH=$archName
  fileName='gosh_'$version'_win_'$archName
  go build -ldflags="-s -w" -o $fileName main.go
  upx --best --no-progress $fileName
  echo "build $archName done"
done
