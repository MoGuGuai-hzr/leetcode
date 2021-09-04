package main

func fib(n int) int {
	f0, f1 := 0, 1
	for i := 0; i < n; i++ {
		f0, f1 = f1, (f0+f1)%(1e9+7)
	}
	return f0 % (1e9 + 7)
}
