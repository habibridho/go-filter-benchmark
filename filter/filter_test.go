package filter

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

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

func BenchmarkFilterWith100kbText(b *testing.B) {
	arr := constructLipsumArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(arr, func(val interface{}) bool {
			return val != "c"
		})
	}
}

func BenchmarkFilterStringWith100kbText(b *testing.B) {
	arr := constructLipsumArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterString(arr, func(val string) bool {
			return val != "c"
		})
	}
}

func constructLipsumArr() []string {
	file, err := os.Open("lipsum.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteArr, _ := ioutil.ReadAll(file)
	arr := strings.Split(string(byteArr), "")
	return arr
}
