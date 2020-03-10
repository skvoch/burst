package model

// PreviewToken - enitiry for with UUID token for uploading the book preview
type PreviewToken struct {
	UID    string `json:"uid"`
	BookID int    `json:"book_id"`
}
