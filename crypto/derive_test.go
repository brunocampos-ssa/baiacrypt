package crypto

import (
	"bytes"
	"testing"
)

func TestDeriveKeyArgon2id(t *testing.T) {
	password := []byte("strong-password")
	salt := []byte("16-byte-salt!!") // must be 16 bytes

	key1 := deriveKeyArgon2id(password, salt)
	key2 := deriveKeyArgon2id(password, salt)

	if len(key1) != 32 {
		t.Errorf("expected key length 32, got %d", len(key1))
	}

	if !bytes.Equal(key1, key2) {
		t.Error("expected derived keys to be equal for same password and salt")
	}

	// Change salt, key should differ
	salt2 := []byte("another-salt-123")
	key3 := deriveKeyArgon2id(password, salt2)
	if bytes.Equal(key1, key3) {
		t.Error("expected different key with different salt")
	}

	// Change password, key should differ
	key4 := deriveKeyArgon2id([]byte("different-password"), salt)
	if bytes.Equal(key1, key4) {
		t.Error("expected different key with different password")
	}
}
