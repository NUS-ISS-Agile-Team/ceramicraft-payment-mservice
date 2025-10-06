package utils

import "math/rand"

const charset = "0123456789"

func GenRedeemCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
