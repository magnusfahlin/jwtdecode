package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <jwt>\n", os.Args[0])
		os.Exit(2)
	}

	token := strings.TrimSpace(os.Args[1])
	parts := strings.Split(token, ".")
	if len(parts) != 3 && len(parts) != 5 {
		fatalf("invalid token: expected 3 parts (JWS) or 5 parts (JWE), got %d", len(parts))
	}

	// Part 0 is always the (protected) header for both JWS and JWE.
	headerJSON, err := decodeB64URLToJSON(parts[0])
	if err != nil {
		fatalf("failed to decode header: %v", err)
	}

	fmt.Println("Header:")
	printPrettyJSON(headerJSON)

	if len(parts) == 3 {
		// JWS: payload is part 1
		payloadJSON, err := decodeB64URLToJSON(parts[1])
		if err != nil {
			fatalf("failed to decode payload: %v", err)
		}

		fmt.Println("\nPayload:")
		printPrettyJSON(payloadJSON)

		// Signature is part 2 (not JSON)
		fmt.Println("\nSignature (base64url):")
		fmt.Println(parts[2])
		return
	}

	// JWE: payload is encrypted (ciphertext is part 3, auth tag is part 4)
	// We can show the raw parts, but can't decode payload without decryption keys.
	fmt.Println("\nJWE parts (payload is encrypted):")
	fmt.Printf("EncryptedKey (base64url): %s\n", parts[1])
	fmt.Printf("IV          (base64url): %s\n", parts[2])
	fmt.Printf("Ciphertext  (base64url): %s\n", parts[3])
	fmt.Printf("AuthTag     (base64url): %s\n", parts[4])
}

func decodeB64URLToJSON(b64url string) (map[string]any, error) {
	if b64url == "" {
		return nil, errors.New("empty base64url segment")
	}

	raw, err := base64.RawURLEncoding.DecodeString(b64url)
	if err != nil {
		return nil, err
	}

	var obj map[string]any
	if err := json.Unmarshal(raw, &obj); err != nil {
		var anyVal any
		if err2 := json.Unmarshal(raw, &anyVal); err2 != nil {
			return nil, fmt.Errorf("not valid JSON: %w", err)
		}
		return map[string]any{"_": anyVal}, nil
	}

	return obj, nil
}

func printPrettyJSON(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("%v\n", v)
		return
	}
	fmt.Println(string(b))
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "jwtdecode: "+format+"\n", args...)
	os.Exit(1)
}
