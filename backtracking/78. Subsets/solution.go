package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	var curr []int
	var explore func(int)
	explore = func(i int) {
		if i == n {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}
		curr = append(curr, nums[i])
		explore(i + 1)
		curr = curr[:len(curr)-1]
		explore(i + 1)
	}
	explore(0)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
