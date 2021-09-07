package invoiceitem

type Item struct {
	id      string
	product string
	value   float64
}

func NewInvoiceItem(id string, product string, value float64) Item {
	return Item{id, product, value}
}
