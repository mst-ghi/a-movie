package helpers

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_KEY", "helpers-package-test-key")
	os.Exit(m.Run())
}
