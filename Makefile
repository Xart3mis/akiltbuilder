release: main.go
	go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -o bin/ .
