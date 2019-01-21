package model

//Wish contains wish's properties
type Wish struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Story string `json:"story"`
}

//WishServicer type contains crud functionalities
type WishServicer interface {
	CreateWish(w *Wish) error
}
