package utils

import "crypto/rand"

const charset = "0123456789"

func GenRedeemCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			// fallback: use '0' if random fails
			code[i] = charset[0]
			continue
		}
		code[i] = charset[int(b[0])%len(charset)]
	}
	return string(code)
}
