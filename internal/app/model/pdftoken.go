package model

// PDFToken - enitiry for with UUID token for uploading the book file
type PDFToken struct {
	UID    string `json:"uid"`
	BookID int    `json:"name"`
}
