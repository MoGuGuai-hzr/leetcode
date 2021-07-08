package main

import "fmt"

// 空间复杂度最低
func numSubarraysWithSum(nums []int, goal int) int {
	cnt := 0
	sum := 0
	left := 0
	for right, v := range nums {
		sum += v

		for left < right && sum > goal {
			sum -= nums[left]
			left++
		}
		if sum == goal {
			cnt++
			tempL := left
			for tempL < right && nums[tempL] == 0 {
				tempL++
				cnt++
			}
		}
	}
	return cnt
}

// 空间换时间
func numSubarraysWithSum2(nums []int, goal int) (ans int) {
	cnt := map[int]int{}
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return
}

func main() {
	for _, test := range []struct {
		nums     []int
		goal, ex int
	}{
		{[]int{1, 0, 1, 0, 1}, 2, 4},
		{[]int{0, 0, 0, 0, 0}, 0, 15},
	} {
		r := numSubarraysWithSum(test.nums, test.goal)
		r2 := numSubarraysWithSum2(test.nums, test.goal)
		fmt.Println(r == test.ex, r2 == test.ex)
	}
}
