package main

import "fmt"

func maximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		t[0] = '2'
		if t[1] != '?' && t[1] >= '4' {
			t[0] = '1'
		}
	}
	if t[1] == '?' {
		t[1] = '3'
		if t[0] != '2' {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}

func main() {
	for _, test := range []struct {
		time, exp string
	}{
		{"2?:?0", "23:50"},
		{"1?:22", "19:22"},
		{"0?:39", "09:39"},
		{"?4:03", "14:03"},
	} {
		r := maximumTime(test.time)
		fmt.Println(r, r == test.exp)
	}
}
