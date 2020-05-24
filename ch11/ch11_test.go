package ch11

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

// 性能剖析,会临时生成ch11.test可执行文件
// go test -run=NONE -bench=BenchmarkIsPalindrome -cpuprofile=cpu.log golang/ch11
// go tool pprof -text -nodecount=10 ./ch11.test cpu.log
// liupeng@bd-dream-liupeng:~/golang/src/golang$ go tool pprof -text -nodecount=10 ./ch11.test cpu.log
// File: ch11.test
// Type: cpu
// Time: May 24, 2020 at 9:39pm (CST)
// Duration: 1.72s, Total samples = 1770ms (102.77%)
// Showing nodes accounting for 1400ms, 79.10% of 1770ms total
// Showing top 10 nodes out of 90
//       flat  flat%   sum%        cum   cum%
//      430ms 24.29% 24.29%     1510ms 85.31%  golang/ch11.IsPalindrome
//      230ms 12.99% 37.29%      430ms 24.29%  runtime.mallocgc
//      190ms 10.73% 48.02%      820ms 46.33%  runtime.growslice
//      160ms  9.04% 57.06%      160ms  9.04%  unicode.ToLower
//      100ms  5.65% 62.71%      100ms  5.65%  unicode.IsLetter
//       90ms  5.08% 67.80%       90ms  5.08%  runtime.memclrNoHeapPointers
//       70ms  3.95% 71.75%       70ms  3.95%  runtime.memmove
//       50ms  2.82% 74.58%       50ms  2.82%  runtime.roundupsize (inline)
//       40ms  2.26% 76.84%       40ms  2.26%  runtime.nextFreeFast (inline)
//       40ms  2.26% 79.10%      130ms  7.34%  runtime.sweepone
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
