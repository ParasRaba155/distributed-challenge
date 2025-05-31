package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/counter"
	"fly.io/distributed-challenge/message"
)

func main() {
	node := maelstrom.NewNode()
	counterHandler := counter.NewCounterHandler(node)
	node.Handle(message.ADD.String(), counterHandler.Handle)
	node.Handle(message.READ.String(), counterHandler.HandleRead)

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
