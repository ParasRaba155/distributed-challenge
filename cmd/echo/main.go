package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/echo"
	"fly.io/distributed-challenge/message"
)

func main() {
	node := maelstrom.NewNode()
	echoHandler := echo.NewEchoHandler(node)
	node.Handle(message.ECHO.String(), echoHandler.Handle)

	err := node.Run()
	if err != nil {
		log.Fatalf("error in running node: %v", err)
	}
}
