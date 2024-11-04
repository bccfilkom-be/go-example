package dto

type Pet struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photoURL"`
	Sold     bool   `json:"sold"`
}
