package main

var (
	x = 0
)

func rand7() int {
	return 0
}

func rand10() int {
	x += rand7()
	return x%10 + 1
}
