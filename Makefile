

fmt:
	gofmt -w *.go

run:
	go run main.go parser.go store.go

test:
	go test -v

