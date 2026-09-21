package abstractfactory

type ecoCup struct{}

func (ecoCup) RimDiameterMM() int {
	return 70
}

func (ecoCup) Material() string {
	return "paperboard"
}

type ecoLid struct{}

func (ecoLid) Fits(cup Cup) bool {
	return cup.RimDiameterMM() == 70
}

func (ecoLid) SealType() string {
	return "snap-fit"
}

type ecoReceipt struct{}

func (ecoReceipt) Render(coffeeName string) string {
	return "Eco receipt for " + coffeeName
}
