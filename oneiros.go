package oneiros

import (
	"crypto/rand"
	"io"
	"math/big"
)

func Generate() string {
	n := new(big.Int)
	bytes := make([]byte, 7)
	for {
		io.ReadFull(rand.Reader, bytes)
		n.SetBytes(bytes)
		if len(n.String()) == 16 {
			return n.String()
		}
	}
}
