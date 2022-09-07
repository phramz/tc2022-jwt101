package signing_algorithms

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func SignRS256() string {
	hasher := sha256.New()
	hasher.Write([]byte(`eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9`)) // header
	hasher.Write([]byte(`.`))                                    // separator
	hasher.Write([]byte(`eyJhdWQiOiJ0ZWNoY2FtcC5oYW1i…`))        // claims

	block, _ := pem.Decode([]byte(`-----BEGIN RSA PRIVATE KEY-----…`))
	rsaKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	signature, _ := rsa.SignPKCS1v15(
		rand.Reader,
		rsaKey,
		crypto.SHA256,
		hasher.Sum(nil))

	return base64.RawURLEncoding.EncodeToString(signature)
}

func VerifyRS256(signature string) error {
	sig, _ := base64.RawURLEncoding.DecodeString(signature)

	hasher := sha256.New()
	hasher.Write([]byte(`eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9`)) // header
	hasher.Write([]byte(`.`))                                    // separator
	hasher.Write([]byte(`eyJhdWQiOiJ0ZWNoY2FtcC5oYW1i…`))        // claims

	block, _ := pem.Decode([]byte(`-----BEGIN PUBLIC KEY-----…`))
	rsaKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	return rsa.VerifyPKCS1v15(rsaKey, crypto.SHA256, hasher.Sum(nil), sig)
}
