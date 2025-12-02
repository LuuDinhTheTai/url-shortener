package encrypt

import (
	"errors"
	"math/big"
	"strings"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const base = 62

// EncodeBase62 encodes a byte slice into a Base62 string
func EncodeBase62(data []byte) string {
	// Convert bytes to a big integer
	num := new(big.Int).SetBytes(data)

	// Special case: zero
	if num.Cmp(big.NewInt(0)) == 0 {
		return string(base62Chars[0])
	}

	var encoded strings.Builder
	for num.Cmp(big.NewInt(0)) > 0 {
		remainder := new(big.Int)
		num.DivMod(num, big.NewInt(base), remainder)
		encoded.WriteByte(base62Chars[remainder.Int64()])
	}

	// Reverse the string because we built it backwards
	runes := []rune(encoded.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// DecodeBase62 decodes a Base62 string into a byte slice
func DecodeBase62(s string) ([]byte, error) {
	num := big.NewInt(0)
	for _, r := range s {
		index := strings.IndexRune(base62Chars, r)
		if index == -1 {
			return nil, errors.New("invalid character in Base62 string")
		}
		num.Mul(num, big.NewInt(base))
		num.Add(num, big.NewInt(int64(index)))
	}
	return num.Bytes(), nil
}
