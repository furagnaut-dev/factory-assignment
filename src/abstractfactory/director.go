package abstractfactory

import "fmt"

type ServingSet struct {
	Cup Cup
	Lid Lid
	ReceiptText string
}

type ServingStation struct {
	factory ServingSetFactory
}

func NewServingStation(factory ServingSetFactory) ServingStation {
	return ServingStation{
		factory: factory,
	}
}

func (station ServingStation) PackageCoffee(coffeeName string) (ServingSet, error) {
	cup := station.factory.CreateCup()
	lid := station.factory.CreateLid()
	receipt := station.factory.CreateReceipt()

	if !lid.Fits(cup) {
		return ServingSet{}, fmt.Errorf(
			"Lid is incompatible with this cup",
		)
	}

	return ServingSet{
		Cup: cup,
		Lid: lid,
		ReceiptText: receipt.Render(coffeeName),
	}, nil
}