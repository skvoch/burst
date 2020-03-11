package apiclient

// BookUploadTokens - this tokens need for upload files
// Send preview to "/v1.0/books/{id}/preview/" endpoint
// Send book to "/v1.0/books/{id}/books/" endpoint
type BookUploadTokens struct {
	BookID      int    `json:"book_id"`
	FileUUID    string `json:"file_uuid"`
	PreviewUUID string `json:"preview_uuid"`
}

// GetBooksIDsResponse - reposnse with books ids
type GetBooksIDsResponse struct {
	BooksIDs []int `json:"books_ids"`
}

// CreateTypeResponse - response with id of created type
type CreateTypeResponse struct {
	ID int `json:"ID"`
}

// FileResponse - response with data of file and meta information
type FileResponse struct {
	Data     []byte
	FileName string
	Err      error
}
