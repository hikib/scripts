# Builds the go binary
podman run --rm -e GOOS=windows -e GOARCH=amd64 -v "./:/app" -w "/app" docker.io/golang go build -o .
