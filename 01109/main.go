package main

func corpFlightBookings1(bookings [][]int, n int) []int {
	rst := make([]int, n)
	for _, list := range bookings {
		start, end, num := list[0], list[1], list[2]
		if start > end {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			rst[i-1] += num
		}
	}
	return rst
}

// 差分
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
