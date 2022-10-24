#! /bin/sh

arch=$(go env GOARCH)

echo "Running test for $arch"

# if amd64, run normally
if [ "$arch" = "amd64" ]; then
    go test --cover ./...
    # if arm, run with tag dynamic
    elif [ "$arch" = "arm64" ]; then
    go test --tags dynamic --cover ./...
else
    echo "Unknown architecture: $arch"
    exit 1
fi