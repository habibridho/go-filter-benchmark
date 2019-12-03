package filter

import "testing"

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []string{"a", "b", "c", "d", "e"}
		Filter(arr, func(val interface{}) bool {
			return val != "c"
		})
	}
}

func BenchmarkFilterString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []string{"a", "b", "c", "d", "e"}
		FilterString(arr, func(val string) bool {
			return val != "c"
		})
	}
}
