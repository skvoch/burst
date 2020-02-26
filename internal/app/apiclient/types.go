package apiclient

// BookUploadTokens - this tokens need for upload files
// Send preview to "/v1.0/books/{id}/preview/" endpoint
// Send book to "/v1.0/books/{id}/books/" endpoint
type BookUploadTokens struct {
	BookID      int    `json:"book_id"`
	FileUUID    string `json:"file_uuid"`
	PreviewUUID string `json:"preview_uuid"`
}

type GetBooksIDsResponse struct {
	BooksIDs []int `json:"books_ids"`
}

type CreateTypeResponse struct {
	ID int `json:"ID"`
}
