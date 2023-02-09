package types

type Orders struct {
	OrderId      int    `json:"OrderId"`
	UserId       int    `json:"UserId"`
	BookId       int    `json:"BookId"`
	Quantity     int    `json:"Quantity"`
	OrderDate    string `json:"OrderDate"`
	PriceOfOrder int    `json:"PriceOfOrder"`
	OrderStatus  string `json:"OrderStatus"`
}
