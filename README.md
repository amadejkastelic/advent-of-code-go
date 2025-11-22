# Advent of Code Solutions in Go

[![go report card](https://goreportcard.com/badge/github.com/amadejkastelic/advent-of-code-go "go report card")](https://goreportcard.com/report/github.com/amadejkastelic/advent-of-code-go)
[![CI status](https://github.com/amadejkastelic/advent-of-code-go/actions/workflows/build.yaml/badge.svg?branch=main "test status")](https://github.com/amadejkastelic/advent-of-code-go/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/amadejkastelic/advent-of-code-go?tab=doc)

## Running

To run for a certain year and day, use the following command:
```bash
› nix run .#default -- -year=2024 -day=1
```

## Development

```bash
› nix develop
```

## Build

```bash
› nix build
```

## Tests

To run all checks, including tests, use:
```bash
› nix flake check
```
