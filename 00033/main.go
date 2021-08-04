package main

func search1(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n
	for i < j {
		middle := (i + j) >> 1
		if nums[middle] >= nums[0] && target >= nums[0] || nums[middle] < nums[0] && target < nums[0] {
			if nums[middle] >= target {
				j = middle
			} else {
				i = middle + 1
			}
		} else if nums[middle] >= nums[0] && target < nums[0] {
			i = middle + 1
		} else {
			j = middle
		}
	}
	if i == n || nums[i] != target {
		return -1
	}
	return i
}

func search(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n
	for i < j {
		middle := (i + j) >> 1
		if nums[middle] >= nums[0] {
			if target >= nums[0] && nums[middle] >= target {
				j = middle
			} else {
				i = middle + 1
			}
		} else {
			if target < nums[0] && nums[middle] < target {
				i = middle + 1
			} else {
				j = middle
			}
		}
	}
	if i == n || nums[i] != target {
		return -1
	}
	return i
}
