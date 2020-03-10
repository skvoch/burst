package store

// Store - this interface aggregate all repositories
type Store interface {
	Books() BooksRepository
	Types() TypesRepository
	TokensPDF() PDFTokenRepository
	TokensPreview() PreviewTokenRepository
}
