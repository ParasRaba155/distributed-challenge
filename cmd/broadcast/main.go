package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/broadcast"
	"fly.io/distributed-challenge/message"
)

func main() {
	node := maelstrom.NewNode()
	broadcastHandler := broadcast.NewBroadcastHandler(node)
	node.Handle(message.BROADCAST.String(), broadcastHandler.HandleBroadcast)
	node.Handle(message.TOPOLOGY.String(), broadcastHandler.HandleTopology)
	node.Handle(message.READ.String(), broadcastHandler.HandleRead)

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
