package factorymethod

type Espresso struct {
}

func (e Espresso) Name() string {
	return "Espresso"
}

func (e Espresso) VolumeML() int {
	return 60
}

func (e Espresso) Prepare() []PreparationStep {
	return []PreparationStep {
		{Action: "Grinding coffee beans..."},
		{Action: "Extract 60 ml of Espresso."},
	}
}

type Americano struct {
}

func (a Americano) Name() string {
	return "Americano"
}

func (a Americano) VolumeML() int {
	return 200
}

func (a Americano) Prepare() []PreparationStep {
	return []PreparationStep {
		{Action: "Grinding coffee beans..."},
		{Action: "Extract 60 ml of Espresso."},
		{Action: "Adding hot water to reach 200 ml of Americano."},
	}
}

type Latte struct{}

func (l Latte) Name() string {
	return "Latte"
}

func (l Latte) VolumeML() int {
	return 200
}

func (l Latte) Prepare() []PreparationStep {
	return []PreparationStep {
		{Action: "Grinding coffee beans..."},
		{Action: "Extract 60 ml of Espresso."},
		{Action: "Adding milk to reach 200 ml of Latte."},
	}
}

var (
	_ Coffee = Espresso{}
	_ Coffee = Americano{}
	_ Coffee = Latte{}
)