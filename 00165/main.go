package main

func compareVersion(version1 string, version2 string) int {
	var i, j int
	for value1, value2 := 0, 0; i < len(version1) && j < len(version2); {
		value1, i = searchValue(version1, i)
		value2, j = searchValue(version2, j)
		if value1 > value2 {
			return 1
		}
		if value1 < value2 {
			return -1
		}
	}
	for i < len(version1) {
		if version1[i] != '.' && version1[i] > '0' {
			return 1
		}
		i++
	}
	for j < len(version2) {
		if version2[j] != '.' && version2[j] > '0' {
			return -1
		}
		j++
	}
	return 0
}

func searchValue(s string, start int) (int, int) {
	value := 0
	var i int
	for i = start; i < len(s); i++ {
		if s[i] == '.' {
			return value, i + 1
		}
		value = value*10 + int(s[i]-'0')
	}
	return value, i
}
