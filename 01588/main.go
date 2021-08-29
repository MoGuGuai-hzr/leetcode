package main

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	n := len(arr)
	var leftCount int
	var rightCount int
	var leftOdd int
	var rightOdd int
	var leftEven int
	var rightEven int
	for i := 0; i < n; i++ {
		leftCount, rightCount = i, n-i-1
		leftOdd = (leftCount + 1) >> 1
		rightOdd = (rightCount + 1) >> 1
		leftEven = leftCount>>1 + 1
		rightEven = rightCount>>1 + 1
		sum += arr[i] * (leftOdd*rightOdd + leftEven*rightEven)
	}
	return sum
}
