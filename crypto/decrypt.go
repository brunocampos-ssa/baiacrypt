package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// decrypt decrypts the given data using AES-GCM with a password.
// The data must be in the format: salt | nonce | ciphertext.
// The salt and nonce are generated during encryption and are required for decryption.
// The password is used to derive the encryption key.
// Returns the decrypted plaintext or an error if decryption fails.
func Decrypt(data []byte, password []byte) ([]byte, error) {
	if len(data) < saltSize+nonceSize {
		return nil, errors.New("invalid encrypted data")
	}
	salt := data[:saltSize]
	nonce := data[saltSize : saltSize+nonceSize]
	ciphertext := data[saltSize+nonceSize:]

	key := deriveKeyArgon2id(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}
