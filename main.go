package main

import(
  "log"
  "github.com/smowafy/caas/cruntime"
)

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
