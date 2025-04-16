package crypto

// This file contains constants used for cryptographic operations.
// These constants are used for key derivation, encryption, and decryption.
const (
	saltSize     = 16
	keySize      = 32 // AES-256
	nonceSize    = 12
	argonTime    = 1
	argonMem     = 64 * 1024 // 64 MB
	argonThreads = 4
)
