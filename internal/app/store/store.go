package store

// Store ...
type Store interface {
	Books() BooksRepository
	Types() TypesRepository
	TokensPDF() PDFTokenRepository
	TokensPreview() PreviewTokenRepository
}
