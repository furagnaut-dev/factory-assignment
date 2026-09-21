package main

import (
	"factory/src/factorymethod"
	"fmt"
)

func main() {
	creator := factorymethod.EspressoCreator{}
	run(creator)
}

func run(creator factorymethod.Creator) {
	coffee := creator.OrderPrepare()

	fmt.Printf("%s - %d ml\n", coffee.Name, coffee.VolumeML)

	for _, step := range coffee.Steps {
		fmt.Println(step.Action)
	}
}