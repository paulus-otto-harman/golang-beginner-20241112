package model

type OrderItem struct {
	BookId   string `json:"bookId"`
	BookName string `json:"bookName,omitempty"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price,omitempty"`
	Subtotal int    `json:"subtotal,omitempty"`
	Discount int    `json:"discount,omitempty"`
}
