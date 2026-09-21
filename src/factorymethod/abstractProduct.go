package factorymethod

type Coffee interface{
	Name() string
	VolumeML() int
	Prepare() []PreparationStep
}

type PreparationStep struct {
	Action string
}

