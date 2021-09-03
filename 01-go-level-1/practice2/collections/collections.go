package collections

type Item int

type Collections interface {
	Add(el Item)
	Delete() Item
	Dump() []Item
}
