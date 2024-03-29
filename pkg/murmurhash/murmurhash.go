package murmurhash

import (
	"fmt"

	. "github.com/rryqszq4/go-murmurhash"
)

func base62Encode(number uint32) string {
	const base = 62
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if number == 0 {
		return string(charset[0])
	}

	encoded := ""
	for number > 0 {
		remainder := number % base
		number /= base
		encoded = string(charset[remainder]) + encoded
	}

	return encoded
}
func GenerateMurmurHash(hashKey string, seed int64) string {
	hash := MurmurHash2([]byte(hashKey), uint32(seed))
	myString := fmt.Sprintf("%v", base62Encode(hash))

	return myString
}
