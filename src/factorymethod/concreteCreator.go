package factorymethod

type EspressoCreator struct {}

func (EspressoCreator) Create() Coffee {
	return Espresso{}
}

func (creator EspressoCreator) OrderPrepare() PreparedCoffee {
	return prepareOrder(creator)
}

type AmericanoCreator struct {}

func (AmericanoCreator) Create() Coffee {
	return Americano{}
}

func (creator AmericanoCreator) OrderPrepare() PreparedCoffee {
	return prepareOrder(creator)
}

type LatteCreator struct {}

func (LatteCreator) Create() Coffee {
	return Latte{}
}

func (creator LatteCreator) OrderPrepare() PreparedCoffee {
	return prepareOrder(creator)
}