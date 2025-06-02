# Distributed System Challenge by [Fly.io](https://fly.io/dist-sys/)

- To get started check the [docs](https://github.com/jepsen-io/maelstrom/blob/main/doc/01-getting-ready/index.md) for installing maelstrom
- Then run the required challenge
- Currently implemented challenges are
    - 1. [Echo](https://fly.io/dist-sys/1/)
    - 2. [Unique ID Generation](https://fly.io/dist-sys/2/)
    - 3. [Broadcast](https://fly.io/dist-sys/3a/)
    - 4. [Grow-Only Counter](https://fly.io/dist-sys/4/)
    - 5. [Single-Node Kafka-Style Log](https://fly.io/dist-sys/5a/)
    - 6. [Single-Node, Totally-Available Transactions](https://fly.io/dist-sys/6a/)
- Code Structure
```bash
├── bin
│   ├── echo
│   └── generate
│   └── <....> # challenges
├── cmd
│   ├── echo
│   │   └── main.go
│   └── generate
│       └── main.go
│   └── <.....> # challenge name
│       └── main.go
├── echo
│   └── echo.go # handler for echo
├── generate
│   └── generate.go # handler for generate
├── <.....> # challenge name
│   └── <name>.go # handler for challenge
├── go.mod
├── go.sum
├── message
│   └── type.go
```
- how to run a challenge
```bash
make test-<challenge-name>
```

## NOTE: In testing [challenge 5: kafka style logs](https://fly.io/dist-sys/5a/) the rate given in docs is 1000, however I kept getting `java.lang.OutOfMemoryError: Java heap space` so I reduced it to 100 One can increase it for their own
