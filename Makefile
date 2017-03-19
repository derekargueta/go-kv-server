

PORT=:8080
PEERS=

fmt:
	gofmt -w *.go

run:
	go run main.go commands.go parser.go store.go $(PORT) $(PEERS)

dbuild:
	docker build -t go-kv-server .

drun:
	docker run \
		--publish 8080:8080 \
		--name test \
		--rm \
		go-kv-server \
		make run

test:
	go test -v

