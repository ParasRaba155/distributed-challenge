package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/kafka"
	"fly.io/distributed-challenge/message"
)

func main() {
	node := maelstrom.NewNode()
	handler := kafka.NewHandler(node)
	node.Handle(message.SEND.String(), handler.Send)
	node.Handle(message.POLL.String(), handler.Poll)
	node.Handle(message.COMMIT_OFFSETS.String(), handler.CommitOffset)
	node.Handle(message.LIST_COMMITTED_OFFSETS.String(), handler.ListCommitedOffset)

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
