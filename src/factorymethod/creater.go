package factorymethod

type PreparedCoffee struct {
	Name string
	VolumeML int
	Steps []PreparationStep
}

type Creator interface {
	Create() Coffee // factory method
	OrderPrepare() PreparedCoffee // business logic
}

func prepareOrder(creator Creator) PreparedCoffee {
	coffee := creator.Create()

	return PreparedCoffee{
		Name: coffee.Name(),
		VolumeML: coffee.VolumeML(),
		Steps: coffee.Prepare(),
	}
}