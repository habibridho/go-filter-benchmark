package main

import (
	"fmt"
	"github.com/habibridho/filter/filter"
)

func main() {
	RunFilterString()
	RunFilterInt()
}

func RunFilterString() {
	arr := []string{"a", "b", "c", "d", "e"}
	result := filter.Filter(arr, func(val interface{}) bool {
		return val != "c"
	})
	fmt.Println(result)
}

func RunFilterInt() {
	arr := []int{1, 2, 3, 4, 5}
	result := filter.Filter(arr, func(val interface{}) bool {
		return val != 5
	})
	fmt.Println(result)
}
