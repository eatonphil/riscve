dist/main: cmd/main.go *.go
	go build -o dist/main $<
