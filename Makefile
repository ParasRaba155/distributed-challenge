MAELSTROM := ~/maelstrom/maelstrom
ECHO_BINARY := ./bin/echo
GENERATE_BINARY := ./bin/generate
BROADCAST_BINARY := ./bin/broadcast
COUNTER_BINARY := ./bin/counter
KAFKA_BINARY := ./bin/kafka
TXN_BINARY := ./bin/txn

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

build-broadcast:
	go build -o bin/broadcast cmd/broadcast/*.go

test-broadcast:build-broadcast
	$(MAELSTROM) test -w broadcast --bin $(BROADCAST_BINARY) --node-count 1 --time-limit 20 --rate 10 

build-counter:
	go build -o bin/counter cmd/counter/*.go

test-counter:build-counter
	$(MAELSTROM) test -w g-counter --bin $(COUNTER_BINARY) --node-count 3 --rate 100 --time-limit 20 --nemesis partition

build-kafka:
	go build -o bin/kafka cmd/kafka/*.go

test-kafka:build-kafka
	$(MAELSTROM) test -w kafka --bin $(KAFKA_BINARY) --node-count 1 --concurrency 2n --time-limit 20 --rate 100

build-txn:
	go build -o bin/txn cmd/txn/*.go

test-txn:build-txn
	$(MAELSTROM) test -w txn-rw-register --bin $(TXN_BINARY) --node-count 1 --time-limit 20 --rate 1000 --concurrency 2n --consistency-models read-uncommitted --availability total
