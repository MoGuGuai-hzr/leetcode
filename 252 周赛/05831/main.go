package main

func numberOfWeeks(milestones []int) int64 {
	max := 0
	sum := 0
	for _, v := range milestones {
		sum += v
		if v > max {
			max = v
		}
	}
	if max > sum-max {
		return 2*int64(sum-max) + 1
	}

	return int64(sum)
}
