package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// f = a_n*26^n + a_{n-1}*26^{n-1} + \dots + a_1*26 + a_0
// a_i \in {1,\dots,26}
// a_i \in {0,\dots,25}
// 所以每次计算 (f-1)%26+1 得到 a_i
// (f-a_i)/26 得到下一项
func convertToTitle1(n int) string {
	b := bytes.NewBuffer(nil)
	w := bufio.NewWriter(b)
	for n > 0 {
		a := (n-1)%26 + 1
		c := 'A' - 1 + a
		w.WriteByte(byte(c))
		n = (n - a) / 26
	}
	w.Flush()
	s := b.Bytes()
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

// 稍微优化一下
func convertToTitle2(n int) string {
	b := bytes.NewBuffer(nil)
	w := bufio.NewWriter(b)
	for n > 0 {
		w.WriteByte(byte((n-1)%26 + 'A'))
		n = (n - 1) / 26
	}
	w.Flush()
	s := b.Bytes()
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(convertToTitle1(1))
	fmt.Println(convertToTitle1(26))
	fmt.Println(convertToTitle1(28))
	fmt.Println(convertToTitle1(701))
	fmt.Println("----")
	fmt.Println(convertToTitle2(1))
	fmt.Println(convertToTitle2(26))
	fmt.Println(convertToTitle2(28))
	fmt.Println(convertToTitle2(701))
}
