package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type string2 []string

func (s *string2) Len() int {
	return len(*s)
}

func (s *string2) Less(i int, j int) bool {
	return strings.Compare((*s)[i], (*s)[j]) < 0
}

func (s *string2) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

type string3 [][]string

func (s *string3) Len() int {
	return len(*s)
}

func (s *string3) Less(i int, j int) bool {
	x, _ := strconv.Atoi((*s)[i][0])
	y, _ := strconv.Atoi((*s)[j][0])
	return x < y
}

func (s *string3) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func addOne(s string) string {
	if s == "" {
		return "1"
	}
	i, _ := strconv.Atoi(s)
	return fmt.Sprint(i + 1)
}

func displayTable(orders [][]string) [][]string {
	// 记录菜品到桌号的映射
	m := make(map[string][]int)
	// 记录桌号
	m1 := make(map[int]bool)
	for i := 0; i < len(orders); i++ {
		// 取到桌号
		tableIndex, _ := strconv.Atoi(orders[i][1])
		// 记录桌号
		m1[tableIndex] = true
		// 把桌号绑定到菜品
		list := m[orders[i][2]]
		list = append(list, tableIndex)
		m[orders[i][2]] = list
	}

	// 拿到所有菜, 排个序
	keyList := make(string2, 0, len(m))
	for key := range m {
		keyList = append(keyList, key)
	}
	sort.Sort(&keyList)
	// 结果大小: 桌子数+1
	tables := make([][]string, 0, len(m1)+1)

	// 首行, 大小为菜品+1
	firstRow := make([]string, 0, len(m)+1)
	// 写入首列
	firstRow = append(firstRow, "Table")
	// 写入其他列
	// 把菜品绑定到列
	m2 := make(map[string]int)
	for i, key := range keyList {
		firstRow = append(firstRow, key)
		m2[key] = i + 1
	}
	tables = append(tables, firstRow)

	// 记录剩余行
	m3 := make(map[int][]string)
	for i := range m1 {
		newRow := make([]string, len(m)+1)
		newRow[0] = fmt.Sprint(i)
		for j := 1; j < len(newRow); j++ {
			newRow[j] = "0"
		}
		tables = append(tables, newRow)
		m3[i] = newRow
	}

	for key, list := range m {
		column := m2[key]
		for _, i := range list {
			row := m3[i]
			row[column] = addOne(row[column])
		}
	}
	s := string3(tables[1:])
	sort.Sort(&s)
	return tables
}

func equal(orders, tables [][]string) bool {
	if len(orders) != len(tables) {
		return false
	}
	for i := 0; i < len(orders); i++ {
		if len(orders[i]) != len(tables[i]) {
			return false
		}
		for j := 0; j < len(orders[i]); j++ {
			if orders[i][j] != tables[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	for i, test := range []struct {
		orders, tables [][]string
	}{
		{
			[][]string{
				{"David", "3", "Ceviche"},
				{"Corina", "10", "Beef Burrito"},
				{"David", "3", "Fried Chicken"},
				{"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"},
			},
			[][]string{
				{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"},
				{"3", "0", "2", "1", "0"},
				{"5", "0", "1", "0", "1"},
				{"10", "1", "0", "0", "0"},
			},
		},
	} {
		r := displayTable(test.orders)
		fmt.Println(i, equal(r, test.tables))
	}
}
