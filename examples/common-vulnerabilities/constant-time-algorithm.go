package common_vulnerabilities

import "crypto/subtle"

func ConstantTimeEqualsBuiltin(x, y []byte) bool {
	return subtle.ConstantTimeCompare(x, y) == 1
}
