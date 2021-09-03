package collections

type Item interface {
}

type Collections interface {
	Add(it Item)
	Delete() Item
	Dump() []Item
}
