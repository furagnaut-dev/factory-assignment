package factorymethod

type Coffee interface{
	VolumeML() int
	Description() string
	Topping() string
	SugarGram() int
	Prepare() []PreparationStep
}