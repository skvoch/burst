package store

import "github.com/skvoch/burst/internal/app/model"

// PDFTokenRepository ...
type PDFTokenRepository interface {
	Create(b *model.PDFToken) error
	GetByUID(UID string) (*model.PDFToken, error)
	Remove(t *model.PDFToken) error
}

// PDFTokenRepository ...
type PreviewTokenRepository interface {
	Create(b *model.PreviewToken) error
	GetByUID(UID string) (*model.PreviewToken, error)
	Remove(t *model.PreviewToken) error
}
