package signing_algorithms

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func SignHS256() string {
	// shared "secret"
	hasher := hmac.New(sha256.New, []byte("secret"))

	// header (base64 encoded)
	hasher.Write([]byte("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9"))
	// separator
	hasher.Write([]byte("."))
	// claims (base64 encoded)
	hasher.Write([]byte("eyJhdWQiOiJ0ZWNoY2FtcC5oYW1idXJnIiwianRpIjoiOTZmNzBk" +
		"ODkiLCJpYXQiOjE2NjEzMzAzMDgsIm5iZiI6MTY2MTMzMDMwOCwiZXhwIjoxNjYxMzM3NT" +
		"A4LCJzdWIiOiJtLnJlaWNoZWwiLCJpc3MiOiJpZC50ZWNoY2FtcC5oYW1idXJnIn0"))

	// returns "mKdydmAO5Mh6bHFBtguwLAdLtxIR3oczRl7hCjsiK0w"
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}
