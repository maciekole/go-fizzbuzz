package main

import (
	"fmt"
	"sort"
)

func fizzbuzz() {
	fbCases := map[int]string{
		3: "fizz",
		5: "buzz",
	}

	keys := make([]int, 0, len(fbCases))

	for k := range fbCases {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 1; i <= 100; i++ {
		printed := false

		for _, key := range keys {
			if i%key == 0 {
				fmt.Print(fbCases[key])
				printed = true
			}
		}
		if !printed {
			fmt.Print(i)
		}

		fmt.Println()
	}
}

func main() {
	fizzbuzz()
}
