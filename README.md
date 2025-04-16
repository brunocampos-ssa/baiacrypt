# 🔐 baiacrypt

**baiacrypt** is a simple and secure CLI tool written in Go for encrypting and decrypting files using modern cryptographic standards.

It combines **AES-256-GCM** for authenticated encryption with **Argon2id** for password-based key derivation, along with a secure password prompt and confirmation during encryption.

## ✨ Features

- 🔐 AES-256-GCM encryption for confidentiality and integrity
- 🧠 Argon2id for strong, memory-hard password-based key derivation
- 👀 Hidden password input via terminal prompt
- ✅ Password confirmation when encrypting
- 📁 Minimal and modular codebase

## 📂 Project Structure

```text
baiacrypt/
├── cmd/
│   └── main.go            # CLI entrypoint (handles encryption/decryption)
├── crypto/
│   ├── crypto.go          # Encryption and decryption logic
│   ├── decrypt.go         # Decryption function
│   ├── derive.go          # Argon2id key derivation
│   └── encrypt.go         # Encryption function
├── prompt/
│   └── password.go        # Secure password prompt and confirmation
├── go.mod
├── go.sum
└── LICENSE                # Apache-2.0 License
```

## 🧰 Build Instructions

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher

### Build from Source

```bash
git clone https://github.com/brunocampos-ssa/baiacrypt.git
cd baiacrypt
go build -o baiacrypt ./cmd
```

This will produce the `baiacrypt` binary in the root folder.

## 🚀 Usage

### 🔒 Encrypt a file

```bash
./baiacrypt encrypt <input_file> <output_file>
```

You will be prompted to enter and confirm a password.

**Example:**
```bash
./baiacrypt encrypt secret.txt secret.enc
```

---

### 🔓 Decrypt a file

```bash
./baiacrypt decrypt <input_file> <output_file>
```

You will be prompted to enter the password used during encryption.

**Example:**
```bash
./baiacrypt decrypt secret.enc secret.txt
```

---

## 🧪 Testing

To run all unit tests:

```bash
go test ./...
```

---

## 🔐 Password Security & Entropy

Although baiacrypt uses top-tier encryption (AES-256-GCM + Argon2id), your **password is the real gatekeeper**. The longer and more complex it is, the harder it becomes to crack.

Assuming users are using **alphanumeric + special characters** (~95 characters set), here’s how different password lengths stack up:

| Password Length | Entropy (bits) | Brute-force Attempts | Time @ 1B guesses/sec | Quantum Resistant? |
|-----------------|----------------|-----------------------|------------------------|---------------------|
| 6               | ~39            | 7.4 × 10⁹             | 7.4 seconds            | ❌ No               |
| 8               | ~52            | 6.6 × 10¹²            | 1.8 hours              | ❌ No               |
| 10              | ~65            | 6.0 × 10¹⁵            | 70 days                | ❌ No               |
| 12              | ~78            | 5.4 × 10¹⁸            | 171 years              | ✅ Yes (for now)    |
| 14              | ~92            | 4.9 × 10²¹            | 1.5 million years      | ✅ Yes              |
| 16              | ~105           | 4.5 × 10²⁴            | 144 billion years      | ✅ Stronger         |
| 20              | ~131           | 3.7 × 10³⁰            | Unbreakable            | ✅ Post-quantum safe |

> 💡 **Recommended:** Use at least a **12-character password** with letters, numbers, and symbols.
> 
> Quantum computers aren't ready for mass attack scenarios **yet**, but strong passwords today will still protect you tomorrow.

---

## 📜 License

This project is licensed under the **Apache License 2.0**.  
See the [LICENSE](./LICENSE) file for more information.

---

> Made with 💻 + ☕ by [@brunocampos-ssa](https://github.com/brunocampos-ssa)
