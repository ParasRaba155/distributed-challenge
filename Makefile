MAELSTROM := /home/paras/maelstrom/maelstrom
BINARY := ./cmd

build:
	go build -o $(BINARY) .

test-echo:build
	$(MAELSTROM) test -w echo --bin $(BINARY) --node-count 1 --time-limit 20

