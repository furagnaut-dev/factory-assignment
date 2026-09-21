package abstractfactory

type EcoFactory struct {}

func (EcoFactory) CreateCup() Cup {
	return ecoCup{}
}

func (EcoFactory) CreateLid() Lid {
	return ecoLid{}
}

func (EcoFactory) CreateReceipt() Receipt {
	return ecoReceipt{}
}

type PremFactory struct {}

func (PremFactory) CreateCup() Cup {
	return premiumCup{}
}

func (PremFactory) CreateLid() Lid {
	return premiumLid{}
}

func (PremFactory) CreateReceipt() Receipt {
	return premiumReceipt{}
}