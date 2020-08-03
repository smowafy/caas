package main

import (
	"github.com/smowafy/caas/cruntime"
	"log"
)

const ContainerId = "my-first-container"
const Container2Id = "my-second-container"

func main() {
	cr, err := cruntime.SetupContainerd()

	if err != nil {
		log.Fatal(err)
	}

	containers, err := cr.ListContainers()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", containers)
}
