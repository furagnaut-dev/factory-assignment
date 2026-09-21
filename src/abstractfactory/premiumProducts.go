package abstractfactory

type premiumCup struct{}

func (premiumCup) RimDiameterMM() int {
	return 80
}

func (premiumCup) Material() string {
	return "double-wall paperboard"
}

type premiumLid struct{}

func (premiumLid) Fits(cup Cup) bool {
	return cup.RimDiameterMM() == 80
}

func (premiumLid) SealType() string {
	return "locking sip seal"
}

type premiumReceipt struct{}

func (premiumReceipt) Render(coffeeName string) string {
	return "Premium receipt for " + coffeeName
}

var (
	_ Cup     = premiumCup{}
	_ Lid     = premiumLid{}
	_ Receipt = premiumReceipt{}
)