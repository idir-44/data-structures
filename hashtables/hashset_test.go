package hashtables

import "testing"

func BenchmarkHashset(b *testing.B) {
	set := Constructor()

	for i := 0; i < 1000001; i++ {
		set.Add(i)
	}
}
