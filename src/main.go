package main

import (
	"fmt"
	"log"

	"factory/src/abstractfactory"
	"factory/src/factorymethod"
)

func main() {
	creator := factorymethod.EspressoCreator{}
	coffee := run(creator)

	var factory abstractfactory.ServingSetFactory =
		abstractfactory.EcoFactory{}

	station := abstractfactory.NewServingStation(factory)

	servingSet, err := station.PackageCoffee(coffee.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(servingSet.Cup.Material())
	fmt.Println(servingSet.Lid.SealType())
	fmt.Println(servingSet.ReceiptText)
}

func run(creator factorymethod.Creator) factorymethod.PreparedCoffee {
	coffee := creator.OrderPrepare()

	fmt.Printf("%s - %d ml\n", coffee.Name, coffee.VolumeML)

	for _, step := range coffee.Steps {
		fmt.Println(step.Action)
	}

	return coffee
}