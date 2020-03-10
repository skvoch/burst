package model

// Type - entity with information of type
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}
