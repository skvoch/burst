package store

import "github.com/skvoch/burst/internal/app/model"

// PDFTokenRepository - provide access to manipulating with tokens for uploading files
type PDFTokenRepository interface {
	Create(b *model.PDFToken) error
	GetByUID(UID string) (*model.PDFToken, error)
	Remove(t *model.PDFToken) error
	RemoveAll() error
}

// PreviewTokenRepository - provide access to manipulating with tokens for uploading previews
type PreviewTokenRepository interface {
	Create(b *model.PreviewToken) error
	GetByUID(UID string) (*model.PreviewToken, error)
	Remove(t *model.PreviewToken) error
	RemoveAll() error
}
