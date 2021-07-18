package main

import (
	"fmt"
	"sort"
)

type sortString []byte

func (s *sortString) Len() int {
	return len(*s)
}

func (s *sortString) Less(i int, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *sortString) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		s := sortString(v)
		sort.Sort(&s)
		list := m[string(s)]
		list = append(list, v)
		m[string(s)] = list
	}
	rst := make([][]string, 0, len(m))
	for _, list := range m {
		rst = append(rst, list)
	}
	return rst
}

func main() {
	for _, test := range []struct {
		strs []string
		exp  [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	} {
		r := groupAnagrams(test.strs)
		fmt.Println(r)
	}
}
