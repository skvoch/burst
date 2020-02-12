package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal()
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", strings.Join(tables, ", "))
			
			if _, err := s.db.Exec(query); err != nil {
				t.Fatal()
			}
		}

		s.Close()
	}
}
