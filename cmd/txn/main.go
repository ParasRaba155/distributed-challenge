package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/message"
	"fly.io/distributed-challenge/txn"
)

func main() {
	node := maelstrom.NewNode()
	handler := txn.NewHandler(node)
	node.Handle(message.TXN.String(), handler.Handle)

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
