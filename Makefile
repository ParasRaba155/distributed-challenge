MAELSTROM := ~/maelstrom/maelstrom
ECHO_BINARY := ./bin/echo
GENERATE_BINARY := ./bin/generate

serve:
	$(MAELSTROM) serve

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

build-echo:
	go build -o bin/echo cmd/echo/*.go

test-echo:build-echo
	$(MAELSTROM) test -w echo --bin $(ECHO_BINARY) --node-count 1 --time-limit 10

build-generate:
	go build -o bin/generate cmd/generate/*.go

test-generate:build-generate
	$(MAELSTROM) test -w unique-ids --bin $(GENERATE_BINARY) --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
