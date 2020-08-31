go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct"

# Set environment variable allow bypassing the proxy for selected modules (optional)
go env -w GOPRIVATE="*.corp.example.com"

set CurrentFolder=%~dp0
cd src/github.com/sonkwo/mmkr
go mod download
cd %CurrentFolder%