#! /bin/sh

cmd=$1
arch=$(go env GOARCH)

echo "Running $cmd for $arch"

# if amd64, run normally
if [ "$arch" = "amd64" ]; then
    go run ./cmd/"$cmd"
    # if arm, run with tag dynamic
    elif [ "$arch" = "arm64" ]; then
    go run --tags dynamic ./cmd/"$cmd"
else
    echo "Unknown architecture: $arch"
    exit 1
fi