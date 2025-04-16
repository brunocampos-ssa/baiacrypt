package crypto

import "golang.org/x/crypto/argon2"

// deriveKeyArgon2id derives a key using Argon2id.
// It takes a password and a salt as input and returns the derived key.
func deriveKeyArgon2id(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, argonTime, argonMem, argonThreads, keySize)
}
