package teststore

import "github.com/skvoch/burst/internal/app/model"

type PDFTokenRepository struct {
	tokens map[string]*model.PDFToken
}

// Create ...
func (p *PDFTokenRepository) Create(token *model.PDFToken) error {

	p.tokens[token.UID] = token

	return nil
}

func (p *PDFTokenRepository) RemoveAll() error {

	p.tokens = make(map[string]*model.PDFToken)

	return nil
}

// GetByUID
func (p *PDFTokenRepository) GetByUID(uid string) (*model.PDFToken, error) {
	return p.tokens[uid], nil
}

// Remove
func (p *PDFTokenRepository) Remove(token *model.PDFToken) error {
	p.tokens[token.UID] = nil

	return nil
}

type PreviewTokenRepository struct {
	tokens map[string]*model.PreviewToken
}

// Create ...
func (p *PreviewTokenRepository) Create(token *model.PreviewToken) error {

	p.tokens[token.UID] = token

	return nil
}

// GetByUID
func (p *PreviewTokenRepository) GetByUID(uid string) (*model.PreviewToken, error) {
	return p.tokens[uid], nil
}

// Remove
func (p *PreviewTokenRepository) Remove(token *model.PreviewToken) error {
	p.tokens[token.UID] = nil

	return nil
}

func (p *PreviewTokenRepository) RemoveAll() error {

	p.tokens = make(map[string]*model.PreviewToken)

	return nil
}
