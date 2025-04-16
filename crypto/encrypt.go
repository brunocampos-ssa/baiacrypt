package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// encrypt encrypts the given plaintext using AES-GCM with a password.
// It generates a random salt and nonce, derives the encryption key using Argon2id,
// and returns the encrypted data in the format: salt | nonce | ciphertext.
// The salt and nonce are required for decryption.
// The password is used to derive the encryption key.
// Returns the encrypted data or an error if encryption fails.
// The salt and nonce are generated randomly and prepended to the ciphertext.
// The salt is used for key derivation, and the nonce is used for AES-GCM encryption.
func Encrypt(plaintext []byte, password []byte) ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	key := deriveKeyArgon2id(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// output = salt | nonce | ciphertext
	return append(append(salt, nonce...), ciphertext...), nil
}
