package abstractfactory

type ServingSetFactory interface {
	CreateCup() Cup
	CreateLid() Lid
	CreateReceipt() Receipt
}