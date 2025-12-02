package encrypt

import (
	"fmt"
	"math/big"
)

type UUID = [16]byte

func EncodeBase62(uuid UUID) string {
	var i big.Int
	i.SetBytes(uuid[:])
	return i.Text(62)
}

func ParseBase62(s string) (UUID, error) {
	var i big.Int
	_, ok := i.SetString(s, 62)
	if !ok {
		return UUID{}, fmt.Errorf("cannot parse base62: %q", s)
	}

	var uuid UUID
	copy(uuid[:], i.Bytes())
	return uuid, nil
}
