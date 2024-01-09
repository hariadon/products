package products

type Item interface {
	ApplyDiscount(discountPercent float64)
}

func ProcessItem(item Item) {
	item.ApplyDiscount(21.5)
}
