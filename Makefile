build:
	go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -o bin/ .