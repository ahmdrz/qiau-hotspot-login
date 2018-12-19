GOCMD=go
GOBUILD=$(GOCMD) build

all: linux darwin windows

linux: deps
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o qiau-login-linux main.go

darwin: deps
		GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o qiau-login-darwin main.go

windows: deps
		GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o qiau-login-windows.exe main.go

deps:
		dep ensure
