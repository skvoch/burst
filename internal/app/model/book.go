package model

// Book ...
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Review      string `json:"review"`
	Rating      int    `json:"rating"`
	Type        int    `json:"type"`
}
