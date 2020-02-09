package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=burst_test sslmode=disable user=postgres database_password = password=docker"
	}

	os.Exit(m.Run())
}
