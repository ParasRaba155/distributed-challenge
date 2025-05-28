package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	node.Handle("echo", EchoHandler(node))

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
