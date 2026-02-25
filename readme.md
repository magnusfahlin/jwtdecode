# jwtdecode

jwtdecode is a simple Go command-line tool that decodes a JWT without validating its signature.

It prints:

- Decoded header
- Decoded payload (for JWS)
- Raw signature
- Raw JWE parts (if the token is encrypted)

> **Important:**
> This tool does not validate signatures, expiration, issuer, audience, or any claims.
> It is intended for debugging and inspection only.

## Installation

### Option 1 – Build locally

```sh
git clone <your-repo-url>
cd jwtdecode
go build -o jwtdecode .
```

### Option 2 – Install with Go

```sh
go install
```

Make sure your `$GOBIN` or `$GOPATH/bin` is in your `PATH`.

## Usage

```sh
jwtdecode <jwt>
```

### Example

```sh
jwtdecode eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Example Output (JWS)

**Header:**
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

**Payload:**
```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```

**Signature (base64url):**
```
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

#### Example Output (JWE)

For encrypted tokens (5 parts), the payload cannot be decoded without keys:

**Header:**
```json
{
  "alg": "RSA-OAEP",
  "enc": "A256GCM"
}
```

**JWE parts (payload is encrypted):**
- EncryptedKey (base64url): ...
- IV          (base64url): ...
- Ciphertext  (base64url): ...
- AuthTag     (base64url): ...

## What It Does

- Splits the token on `.`
- Base64URL decodes header and payload
- Pretty-prints JSON
- Works for:
  - JWS (3 parts)
  - JWE (5 parts)

## What It Does NOT Do

- ❌ Signature validation
- ❌ Expiration checks
- ❌ Issuer / audience validation
- ❌ Decryption of JWE payload
- ❌ Key handling

## Requirements

- Go 1.20+ (or newer)

## Why?

When debugging auth issues, you often just want to quickly inspect what’s inside a token without:

- importing a heavy JWT library
- verifying signatures
- wiring up keys

jwtdecode keeps it minimal and dependency-free.

## License

This project is licensed under the MIT License.

See the LICENSE file for details.