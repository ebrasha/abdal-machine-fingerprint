<div align="right">
  <img src="shot.webp" alt="Abdal Machine Fingerprint" />
</div>

**English** | [فارسی](README.fa.md)

# Abdal Machine Fingerprint

Cross-platform Go library for stable machine identity and cryptographically derived machine fingerprints.

**Module:** `github.com/ebrasha/abdal-machine-fingerprint`  
**Package:** `abdalmf`  
**Version:** `1.0.0` (`abdalmf.Version`)  
**Go:** 1.25+

---

## 🎯 Why This Project Exists

Many applications need a **stable machine identity** for licensing, trial management, device binding, audit trails, or security controls — without relying on fragile hardware serial numbers or invasive fingerprinting.

**Abdal Machine Fingerprint** solves this by reading the operating system's own machine identifier (Windows `MachineGuid`, Linux `machine-id`, macOS `IOPlatformUUID`) and exposing a clean API to hash or protect that identity using industry-standard cryptography.

It is designed for:

- Desktop and workstation software
- Dedicated servers
- VPS and cloud instances
- Offline, privacy-conscious environments

The library never performs network requests, telemetry, or automatic logging.

---

## ✨ Features

- **Raw Machine ID** — `AbdalID()` returns the normalized OS machine identifier
- **Hashed Fingerprints** — `AbdalHash()` with 16 selectable hash algorithms
- **Protected Fingerprints** — `AbdalProtected()` with HMAC or keyed BLAKE2
- **Explicit HMAC API** — `AbdalHMAC()` for application-specific fingerprints
- **Custom Encoding** — `AbdalEncode()` supports HEX, HEX-UPPER, BASE64, BASE64URL
- **Cross-Platform** — Windows, Linux, macOS
- **VPS / Cloud Friendly** — uses OS-level identity, not motherboard or CPU serials
- **Fully Offline** — no HTTP, telemetry, analytics, or tracking
- **Thread-Safe** — safe for concurrent use from multiple goroutines
- **Deterministic** — same inputs always produce the same output

---

## 🚀 How to Use

### Installation

```bash
go get github.com/ebrasha/abdal-machine-fingerprint
```

### Import

```go
import "github.com/ebrasha/abdal-machine-fingerprint"
```

### Raw Machine ID

```go
id, err := abdalmf.AbdalID()
if err != nil {
    // handle error
}
```

> **Privacy:** `AbdalID()` returns the raw machine identifier. Storage, display, and transmission are the developer's responsibility.

### SHA-256 Hash

```go
fingerprint, err := abdalmf.AbdalHash("SHA256")
```

### SHA-512 Hash

```go
fingerprint, err := abdalmf.AbdalHash("SHA512")
```

### SHA3-512 Hash

```go
fingerprint, err := abdalmf.AbdalHash("SHA3-512")
```

### Protected Fingerprint

```go
fingerprint, err := abdalmf.AbdalProtected(
    "abdal-security-tools",
    "SHA256",
)
```

### HMAC Fingerprint

```go
fingerprint, err := abdalmf.AbdalHMAC(
    "abdal-security-tools",
    "SHA512",
)
```

### Custom Output Encoding

```go
fingerprint, err := abdalmf.AbdalEncode(
    "abdal-security-tools",
    "SHA256",
    "BASE64URL",
)
```

---

## 🔐 Supported Hash Algorithms

`MD5`, `SHA1`, `SHA224`, `SHA256`, `SHA384`, `SHA512`, `SHA512-224`, `SHA512-256`, `SHA3-224`, `SHA3-256`, `SHA3-384`, `SHA3-512`, `BLAKE2B-256`, `BLAKE2B-384`, `BLAKE2B-512`, `BLAKE2S-256`

Algorithm names are **case-insensitive** and **format-tolerant** (`SHA-256`, `sha256`, `SHA_256`, `SHA3512`).

Default output for hash and protected APIs is **lowercase hexadecimal**.

---

## 📦 Supported Encodings

| Encoding | Description |
|----------|-------------|
| `HEX` | Lowercase hex (default) |
| `HEX-UPPER` | Uppercase hex |
| `BASE64` | Standard Base64 |
| `BASE64URL` | URL-safe Base64 without padding |

---

## 🖥️ Platform Sources

| Platform | Source |
|----------|--------|
| Windows | `HKLM\SOFTWARE\Microsoft\Cryptography\MachineGuid` |
| Linux | `/var/lib/dbus/machine-id` → `/etc/machine-id` |
| macOS | `IOPlatformUUID` via native IOKit (CGO) |

### macOS Build Note

macOS uses native IOKit via CGO. Build with CGO enabled:

```bash
CGO_ENABLED=1 go build ./...
```

---

## ⚠️ Error Handling

Use `errors.Is()` with sentinel errors:

| Error | When |
|-------|------|
| `abdalmf.ErrMachineIDNotFound` | Machine ID unavailable or empty |
| `abdalmf.ErrUnsupportedPlatform` | OS not supported |
| `abdalmf.ErrUnsupportedAlgorithm` | Unknown hash algorithm |
| `abdalmf.ErrUnsupportedEncoding` | Unknown output encoding |
| `abdalmf.ErrEmptyApplicationID` | Empty application ID in protected calls |
| `abdalmf.ErrUnsupportedHMACAlgorithm` | Algorithm not valid for HMAC/protected |

---

## 📁 Project Structure

```
abdal-machine-fingerprint/
├── abdalconstants.go
├── id.go, hash.go, protected.go, encode.go
└── internal/
    ├── encode/
    ├── hash/
    ├── machine/
    ├── normalize/
    └── platform/
```

Access runtime version:

```go
abdalmf.Version
```

---

## 🐛 Reporting Issues

If you encounter any issues or have configuration problems, please reach out via email at Prof.Shafiei@Gmail.com. You can also report issues on GitLab or GitHub.

## ❤️ Donation

If you find this project helpful and would like to support further development, please consider making a donation:

- [Donate Here](https://t.me/AbdalDonationBot)

## 🤵 Programmer

Handcrafted with Passion by **Ebrahim Shafiei (EbraSha)**

- **E-Mail**: Prof.Shafiei@Gmail.com
- **Telegram**: [@ProfShafiei](https://t.me/ProfShafiei)

## 📜 License

This project is licensed under the AGPLv3 License.
