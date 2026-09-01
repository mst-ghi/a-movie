package helpers

import (
	"testing"

	"app/core/config"
)

func TestGetAppKey_Stable(t *testing.T) {
	t.Setenv("APP_KEY", "a-test-secret")

	first := config.GetAppKey()
	second := config.GetAppKey()

	if len(first) != 32 {
		t.Fatalf("expected 32-byte AES-256 key, got %d bytes", len(first))
	}
	if string(first) != string(second) {
		t.Fatal("key derived from the same APP_KEY must be stable")
	}
}

func TestEncryptDecryptRoundTrip(t *testing.T) {
	t.Setenv("APP_KEY", "a-test-secret")

	plaintext := "hello a-movie"

	encrypted, err := Encrypt(plaintext)
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Fatalf("decrypt failed: %v", err)
	}

	if decrypted != plaintext {
		t.Fatalf("round trip mismatch: got %q, want %q", decrypted, plaintext)
	}
}

func TestDecrypt_RejectsTamperedInput(t *testing.T) {
	t.Setenv("APP_KEY", "a-test-secret")

	encrypted, err := Encrypt("hello")
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}

	if _, err := Decrypt(encrypted + "x"); err == nil {
		t.Fatal("expected error for tampered ciphertext")
	}
}
