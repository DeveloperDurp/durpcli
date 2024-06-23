package main

import (
	"log"

	"gitlab.com/DeveloperDurp/DurpCLI/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
