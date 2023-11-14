package main

import (
	"log"

	"github.com/mfbmina/enxame/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
