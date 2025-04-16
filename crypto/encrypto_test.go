package crypto

import (
	"bytes"
	"testing"
)

func TestEncryptAndDecrypt_Success(t *testing.T) {
	password := []byte("unit-test-password")
	plaintext := []byte("this is a top secret message")

	ciphertext, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	if bytes.Contains(ciphertext, plaintext) {
		t.Error("plaintext should not be visible in ciphertext")
	}

	decrypted, err := Decrypt(ciphertext, password)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if !bytes.Equal(decrypted, plaintext) {
		t.Error("decrypted message does not match original")
	}
}

func TestDecrypt_WithWrongPassword(t *testing.T) {
	plaintext := []byte("this data must be protected")
	password := []byte("correct-password")
	wrongPassword := []byte("wrong-password")

	ciphertext, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	decrypted, err := Decrypt(ciphertext, wrongPassword)
	if err == nil {
		t.Error("expected decryption to fail with wrong password")
	}
	if decrypted != nil {
		t.Errorf("expected nil result with wrong password, got: %s", decrypted)
	}
}

func TestDecrypt_WithTamperedCiphertext(t *testing.T) {
	password := []byte("secure-password")
	plaintext := []byte("authenticated data")

	ciphertext, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	// Tamper with the ciphertext
	ciphertext[15] ^= 0xFF

	_, err = Decrypt(ciphertext, password)
	if err == nil {
		t.Error("expected decryption to fail due to tampered ciphertext")
	}
}
