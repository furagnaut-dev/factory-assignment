package abstractfactory

type Cup interface {
	RimDiameterMM() int
	Material() string
}

type Lid interface {
	Fits(cup Cup) bool
	SealType() string
}

type Receipt interface {
	Render(coffeeName string) string
}