arch=('amd64' '386' 'arm' 'arm64' 'mipsle')
version="0.0.0"
# shellcheck disable=SC2068
for archName in ${arch[@]}; do
  export GOARCH=$archName
  fileName="gosh-$version-$archName"
  go build -ldflags="-s -w" -o $fileName main.go
  upx --best --no-progress $fileName
  echo "build $archName done"
done
