package main

func checkRecord(s string) bool {
	absent, late := 0, 0
	for _, c := range s {
		switch c {
		case 'A':
			absent++
			late = 0
		case 'L':
			late++
		case 'P':
			late = 0
		}
		if absent >= 2 || late >= 3 {
			return false
		}
	}
	return true
}
