# Distributed System Challenge by [Fly.io](https://fly.io/dist-sys/)

- To get started check the [docs](https://github.com/jepsen-io/maelstrom/blob/main/doc/01-getting-ready/index.md) for installing maelstrom
- Then run the required challenge
- Currently implemented challenges are
    - 1. [Echo](https://fly.io/dist-sys/1/)
    - 2. [Unique ID Generation](https://fly.io/dist-sys/2/)
    - 3. [Broadcast](https://fly.io/dist-sys/3a/)
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
