package main

type TimeMap struct {
	tm map[string][]timeType
}

type timeType struct {
	value     string
	timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{tm: make(map[string][]timeType)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	l := this.tm[key]
	l = append(l, timeType{value: value, timestamp: timestamp})
	this.tm[key] = l
}

func (this *TimeMap) Get(key string, timestamp int) string {
	l := this.tm[key]
	if l == nil || l[0].timestamp > timestamp {
		return ""
	}
	var left, middle int
	right := len(l) - 1
	for left <= right {
		middle = (left + right) >> 1
		if l[middle].timestamp > timestamp {
			right = middle - 1
		} else if l[middle].timestamp < timestamp {
			left = middle + 1
		} else {
			return l[middle].value
		}
	}
	if l[middle].timestamp < timestamp {
		return l[middle].value
	} else {
		return l[middle-1].value
	}
}
