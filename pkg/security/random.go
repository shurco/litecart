package security

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

const (
	DefaultIdLength   = 15
	DefaultIdAlphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func RandomString() string {
	b := make([]byte, DefaultIdLength)
	max := big.NewInt(int64(len(DefaultIdAlphabet)))

	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			// fallback: derive pseudo-random index from time hash
			h := sha256.Sum256([]byte(fmt.Sprintf("%d-%d", time.Now().UnixNano(), i)))
			idx := int(h[i%len(h)]) % len(DefaultIdAlphabet)
			if idx < 0 {
				idx = -idx
			}
			b[i] = DefaultIdAlphabet[idx]
			continue
		}
		b[i] = DefaultIdAlphabet[n.Int64()]
	}

	return string(b)
}
