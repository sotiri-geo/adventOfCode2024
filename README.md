# Advent of code 2024


This is part of a pet project in order to learn Go. Inspired by [Learn Go with tests](https://quii.gitbook.io/learn-go-with-test) blog to apply TDD methodology whilst attempting to solve AOC problems.

Project structure is as follows:

```sh
adventOfCode2024/
│── go.mod
│── go.sum
│── README.md
│── cmd/
│   └── day01/
│       └── main.go          # Runs the solution for day 1
│── internal/
│   └── day01/
│       ├── solve.go         # Part 1 & Part 2 logic
│       └── solve_test.go    # Unit tests for day 1
│── inputs/
│   └── day01.txt            # Your puzzle input
```

To run a `day01` solution

```sh
go run ./cmd/day01
```

To run tests for `day01`

```sh
go test ./internal/day01
```
