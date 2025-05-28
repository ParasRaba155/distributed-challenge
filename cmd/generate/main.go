package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/generate"
	"fly.io/distributed-challenge/message"
)

func main() {
	node := maelstrom.NewNode()
	generateHandler := generate.NewGenerateHanlder(node)
	node.Handle(message.GENERATE.String(), generateHandler.Handle)

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
