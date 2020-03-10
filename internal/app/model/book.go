package model

// Book - entity with information of book
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Review      string `json:"review"`
	Rating      int    `json:"rating"`
	Type        int    `json:"type"`
	FilePath    string `json:"file_path,omitempty"`
	PreviewPath string `json:"preview_path,omitempty"`
}

// Sanitaize - remove private paths
func (b *Book) Sanitaize() {
	b.FilePath = ""
	b.PreviewPath = ""
}
