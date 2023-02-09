package types

type Books struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	BookQuantity int    `json:"bookQuantity"`
}
