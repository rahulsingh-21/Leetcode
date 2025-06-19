package main

type TimeMap struct {
	m map[string][]PairII
}

type PairII struct {
	value string
	ts    int
}

func ConstructorX() TimeMap {
	return TimeMap{map[string][]PairII{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], PairII{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	val := this.m[key]
	ans := ""
	l, r := 0, len(val)-1

	for l <= r {
		mid := l + (r-l)/2
		if val[mid].ts == timestamp {
			return val[mid].value
		} else if val[mid].ts > timestamp {
			r = mid - 1
		} else {
			ans = val[mid].value
			l = mid + 1
		}
	}

	return ans
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
