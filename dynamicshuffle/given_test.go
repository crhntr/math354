package dynamicshuffle_test

import (
	"math/rand"
	"testing"

	. "github.com/crhntr/alg/dynamicshuffle"
)

func TestIsShuffleRecursive_000111(t *testing.T) {
	w, u, v := "010101", "000", "111"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_0(t *testing.T) {
	w, u, v := "0", "", "0"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_1(t *testing.T) {
	w, u, v := "1", "1", ""

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursiveLongRandom(t *testing.T) {
	w, u, v := "010101", "000", "111"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_011011(t *testing.T) {
	w, u, v := "001111", "011", "011"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_101101(t *testing.T) {
	w, u, v := "110101", "101", "101"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaaaccebbbbdde(t *testing.T) {
	w, u, v := "aabbababcdcdee", "aaaacce", "bbbbdde"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaaaccqbbbbdde(t *testing.T) {
	w, u, v := "aabbababcdcdee", "aaaaccq", "bbbbdde"

	if IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaabbbabcabc(t *testing.T) {
	w, u, v := "abcabc", "aaa", "bbb"

	if IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

var (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func benchmarkShuffle(wuv []string, f func(w, u, v string) bool, b *testing.B) {
	// b.Logf("v: %q, u: %q, y: %q\n", wuv[0], wuv[1], wuv[2])
	for n := 0; n < b.N; n++ {
		if !f(wuv[0], wuv[1], wuv[2]) {
			b.Fail()
		}
	}
}

func BenchmarkIsShuffleRecursive6(b *testing.B) {
	benchmarkShuffle(generateShuffle(3, 3, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive12(b *testing.B) {
	benchmarkShuffle(generateShuffle(6, 6, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive120(b *testing.B) {
	benchmarkShuffle(generateShuffle(60, 60, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive1200(b *testing.B) {
	benchmarkShuffle(generateShuffle(600, 600, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive12000(b *testing.B) {
	benchmarkShuffle(generateShuffle(6000, 6000, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive120000(b *testing.B) {
	benchmarkShuffle(generateShuffle(60000, 60000, alphabet), IsShuffleRecursive, b)
}

func generateShuffle(uLen, vLen int, Σ string) []string {
	w, u, v := "", "", ""
	for i := 0; i < uLen; i++ {
		index := rand.Intn(len(Σ))
		u += Σ[index : index+1]
	}

	for i := 0; i < vLen; i++ {
		index := rand.Intn(len(Σ))
		v += Σ[index : index+1]
	}

	for ui, vi := 0, 0; ui+vi < (uLen + vLen); {
		if rand.Intn(2) == 1 {
			u += u[ui : ui+1]
			ui++
		} else {
			v += v[vi : vi+1]
			vi++
		}

		if ui == uLen {
			w += v[vi:]
			break
		}
		if vi == vLen {
			w += u[ui:]
			break
		}
	}

	return []string{w, u, v}
}