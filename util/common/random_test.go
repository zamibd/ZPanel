package common

import (
	"testing"
)

func TestRandom_Length(t *testing.T) {
	for _, n := range []int{0, 1, 8, 16, 64} {
		got := Random(n)
		if len(got) != n {
			t.Errorf("Random(%d) returned string of length %d", n, len(got))
		}
	}
}

func TestRandom_Charset(t *testing.T) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	valid := make(map[rune]bool, len(charset))
	for _, c := range charset {
		valid[c] = true
	}
	s := Random(256)
	for _, c := range s {
		if !valid[c] {
			t.Errorf("Random returned unexpected character: %q", c)
		}
	}
}

func TestRandom_NotAlwaysSame(t *testing.T) {
	// With length 16, the probability of getting the same string twice is negligible.
	a := Random(16)
	b := Random(16)
	if a == b {
		// Try once more before failing — the chance of two collisions is astronomically small.
		c := Random(16)
		if a == c {
			t.Error("Random appears to always return the same string")
		}
	}
}

func TestRandomInt_InRange(t *testing.T) {
	for _, n := range []int{1, 2, 10, 100} {
		for i := 0; i < 100; i++ {
			got := RandomInt(n)
			if got < 0 || got >= n {
				t.Errorf("RandomInt(%d) = %d, out of range [0, %d)", n, got, n)
			}
		}
	}
}
