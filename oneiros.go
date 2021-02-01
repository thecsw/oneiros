package oneiros

import (
	"crypto/rand"
	"io"
	"math/big"
)

var (
	// lower is the lower boundary (inclusive).
	lower = big.NewInt(1000000000000000)
	// upper is the upper boundary (inclusive).
	upper = big.NewInt(9999999999999999)
	// t is discarding garbage values
	t = new(big.Int)
	// bytes is the intermediate buffer to read /dev/urandom.
	bytes = make([]byte, 7)
	// n is the random number container
	n = new(big.Int)
)

// Generate will return a cryptographically secure random 16-digit integer.
func Generate() *big.Int {
	for {
		io.ReadFull(rand.Reader, bytes)
		n.SetBytes(bytes)
		// n >= lower <=> n - lower >= 0 <=> sign != -1
		// n <= upper <=> upper - n >= 0 <=> sign != -1
		if t.Sub(upper, n).Sign() != -1 && t.Sub(n, lower).Sign() != -1 {
			return n
		}
	}
}
