package main

import (
	"livepixelshtmx/internal/cli"
	"log"
)

func main() {

	// capturing flags from user
	var cli *cli.CliConfig = &cli.CliConfig{}
	if err := cli.ParseArgs(); err != nil {
		log.Fatalf("Failed to parse arguments: (%s)", err.Error())
	}

	// spawn

}
