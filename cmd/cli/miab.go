package main

import "github.com/rverst/go-miab/cmd/cli/command"

const version = "1.0.0-beta1"

func main() {
	command.Execute(version)
}
