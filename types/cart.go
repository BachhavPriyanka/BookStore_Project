package types

type Cart struct {
	BookName string `json:bookName`
	BookID   int    `json:"bookID"`
	CartID   int    `json:"cartID"`
	Quantity int    `json:"quantity"`
}
