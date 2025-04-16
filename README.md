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
│   └── derive.go          # Argon2id key derivation
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

## 🧪 Testing (optional)

If/when tests are added:

```bash
go test ./...
```

## 📜 License

This project is licensed under the **Apache License 2.0**.  
See the [LICENSE](./LICENSE) file for more information.

---

> Made with 💻 + ☕ by [@brunocampos-ssa](https://github.com/brunocampos-ssa)
