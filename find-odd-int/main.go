package main

import "fmt"

func main() {
	fmt.Printf("Input: [7]\nOutput: %v\n", FindOdd([]int{7}))
	fmt.Printf("Input: [0]\nOutput: %v\n", FindOdd([]int{0}))
	fmt.Printf("Input: [1,1,2]\nOutput: %v\n", FindOdd([]int{1, 1, 2}))
	fmt.Printf("Input: [0,1,0,1,0]\nOutput: %v\n", FindOdd([]int{0, 1, 0, 1, 0}))
	fmt.Printf("Input: [1,2,2,3,3,3,4,3,3,3,2,2,1]\nOutput: %v\n", FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))

}
func FindOdd(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
		println(res)
	}
	return res
}

// func FindOdd(seq []int) int {
// 	m := make(map[int]int)
// 	for _, n := range seq {
// 		addNumberToMap(n, m)
// 	}
// 	return fOdd(m)
// }

func fOdd(m map[int]int) int {
	for k, v := range m {
		if isOdd(v) {
			return k
		}
	}
	return 0
}

func addNumberToMap(n int, m map[int]int) {
	for k := range m {
		if k == n {
			m[n]++
			return
		}
	}
	m[n] = 1
}

func isOdd(n int) bool {
	return n%2 == 1
}
