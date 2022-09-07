package common_vulnerabilities

import (
	"bytes"
	"crypto/sha256"
	"hash"
	"io"
	"net/http"
)

func Verify(token string) error {
	header, _ := MustDecode(token)
	resp, _ := http.Get(`https://id.silpion.de/.well-known/certs/public-key.pem`)
	body, _ := io.ReadAll(resp.Body)
	key := bytes.NewBuffer(body).String()

	return MustGetVerifierFor(header["alg"].(string), key).Verify(token)
}

func MustGetVerifierFor(algo, key string) Verifier {
	switch algo {
	case "HS256":
		return NewHMACVerifier(key, sha256.New())
	case "RS256":
		return NewRSAVerifier(key, sha256.New())
		// etc, etc...
	}

	panic("unsupported algorithm")
}

type Verifier interface {
	Verify(token string) error
}

func NewHMACVerifier(key string, hash hash.Hash) Verifier {
	panic("not implemented")
}

func NewRSAVerifier(key string, hash hash.Hash) Verifier {
	panic("not implemented")
}
func MustDecode(token string) (header, claims map[string]interface{}) {
	panic("not implemented")
}
