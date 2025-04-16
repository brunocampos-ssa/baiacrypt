package crypto

import (
	"bytes"
	"testing"
)

func TestDecrypt_Success(t *testing.T) {
	password := []byte("strong-pass")
	message := []byte("confidential content")

	encrypted, err := Encrypt(message, password)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if !bytes.Equal(decrypted, message) {
		t.Errorf("expected %s, got %s", message, decrypted)
	}
}

func TestDecrypt_InvalidSaltSize(t *testing.T) {
	password := []byte("secret")
	invalidData := make([]byte, 10) // too short, must be salt + nonce + ciphertext

	_, err := Decrypt(invalidData, password)
	if err == nil {
		t.Error("expected error due to insufficient data (invalid salt/nonce size)")
	}
}

func TestDecrypt_TamperedCiphertext(t *testing.T) {
	password := []byte("integrity-check")
	message := []byte("auth-bound content")

	encrypted, err := Encrypt(message, password)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	encrypted[25] ^= 0x99 // tamper a byte

	_, err = Decrypt(encrypted, password)
	if err == nil {
		t.Error("expected error due to tampered ciphertext")
	}
}

func TestDecrypt_WrongPassword(t *testing.T) {
	password := []byte("correct-pass")
	wrong := []byte("wrong-pass")
	message := []byte("sensitive")

	encrypted, err := Encrypt(message, password)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	_, err = Decrypt(encrypted, wrong)
	if err == nil {
		t.Error("expected error due to wrong password")
	}
}
