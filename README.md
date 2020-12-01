# Advent of Code 2020

[![](https://github.com/nlowe/aoc2020/workflows/CI/badge.svg)](https://github.com/nlowe/aoc2020/actions) [![Coverage Status](https://coveralls.io/repos/github/nlowe/aoc2020/badge.svg?branch=master)](https://coveralls.io/github/nlowe/aoc2020?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/nlowe/aoc2020)](https://goreportcard.com/report/github.com/nlowe/aoc2020) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

Solutions for the 2020 Advent of Code

## Building

This project makes use of Go 1.15.

```bash
go mod download
go test ./...
```

## Running the Solutions

To run a solution, use the problem name followed by the path to an input file.

For example, to run problem 2a:

```bash
$ go run ./main.go 2a ./day2/input.txt
Answer: 9633
Took 999.4µs
```

## Adding New Solutions

A generator program is included that makes templates for each day, automatically
downloading challenge input and updating the root command to add new subcommands
for each problem. Running `go generate` from the repo root will generate the
following for each day that is currently accessible:

* `challenge/day<N>/import.go`: A "glue" file combining commands for both of the day's problems to simplify wiring up subcommands
* `challenge/day<N>/a.go`: The main problem implementation, containing a cobra command `A` and the implementation `func a(*challenge.Input) int`
* `challenge/day<N>/a_test.go`: A basic test template
* `challenge/day<N>/input.txt`: The challenge input

Additionally, `challenge/cmd/cmd.go` will be regenerated to import and add all
subcommands.

This requires `goimports` be available on your `$PATH`. Additionally, you must be
logged into https://adventofcode.com in Chrome so the generator can use your session
cookie to download challenge input.

Existing solutions and challenge inputs will be skipped instead of regenerated.

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.
