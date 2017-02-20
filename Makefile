

PORT=:8080
PEERS=

fmt:
	gofmt -w *.go

run:
	go run main.go commands.go parser.go store.go $(PORT) $(PEERS)

test:
	go test -v

