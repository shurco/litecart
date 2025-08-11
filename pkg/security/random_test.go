package security

import (
	"testing"
)

func TestRandomString_LengthAndAlphabet(t *testing.T) {
	const iterations = 100
	for i := 0; i < iterations; i++ {
		s := RandomString()
		if len(s) != DefaultIdLength {
			t.Fatalf("unexpected length: got %d want %d", len(s), DefaultIdLength)
		}
		for _, r := range s {
			if r < '0' || (r > '9' && r < 'a') || r > 'z' { // alphabet is [a-z0-9]
				t.Fatalf("unexpected rune %q in %q", r, s)
			}
		}
	}
}
