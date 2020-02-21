package psqlstore

import "github.com/skvoch/burst/internal/app/model"

type PDFTokenRepository struct {
	store *Store
}

// Create ...
func (p *PDFTokenRepository) Create(token *model.PDFToken) error {
	if err := p.store.db.QueryRow(
		"INSERT INTO pdf_tokens (uid, bookID) VALUES ($1, $2) RETURNING uid ",
		token.UID,
	).Scan(token.UID); err != nil {
		return err
	}

	return nil
}

// GetByUID
func (p *PDFTokenRepository) GetByUID(uid string) (*model.PDFToken, error) {
	token := &model.PDFToken{}

	err := p.store.db.QueryRow(
		"SELECT uid, bookID FROM pdf_tokens WHERE uid = $1;",
		uid,
	).Scan(&token.UID, &token.BookID)

	if err != nil {
		return nil, err
	}

	return token, nil
}

// Remove
func (p *PDFTokenRepository) Remove(token *model.PDFToken) error {

	return p.store.db.QueryRow("DELETE FROM pdf_tokens WHERE uid = $1;", token.UID).Scan()
}

// RemoveAll ...
func (p *PDFTokenRepository) RemoveAll() error {
	if _, err := p.store.db.Query("TRUNCATE TABLE pdf_tokens CASCADE;"); err != nil {
		return err
	}

	return nil
}

type PreviewTokenRepository struct {
	store *Store
}

// Create ...
func (p *PreviewTokenRepository) Create(token *model.PreviewToken) error {
	if err := p.store.db.QueryRow(
		"INSERT INTO preview_tokens (uid, bookID) VALUES ($1, $2) RETURNING uid",
		token.UID,
		token.BookID,
	).Scan(token.UID); err != nil {
		return err
	}

	return nil
}

// GetByUID
func (p *PreviewTokenRepository) GetByUID(uid string) (*model.PreviewToken, error) {
	token := &model.PreviewToken{}

	err := p.store.db.QueryRow(
		"SELECT uid, bookID FROM preview_tokens WHERE uid = $1;",
		uid,
	).Scan(&token.UID, &token.BookID)

	if err != nil {
		return nil, err
	}

	return token, nil
}

// Remove
func (p *PreviewTokenRepository) Remove(token *model.PreviewToken) error {

	return p.store.db.QueryRow("DELETE FROM preview_tokens WHERE uid = $1;", token.UID).Scan()
}

// RemoveAll ...
func (p *PreviewTokenRepository) RemoveAll() error {
	if _, err := p.store.db.Query("RUNCATE TABLE preview_tokens CASCADE;"); err != nil {
		return err
	}

	return nil
}
