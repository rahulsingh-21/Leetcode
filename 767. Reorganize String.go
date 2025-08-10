package main

import "container/heap"

func reorganizeString(s string) string {
	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	h := &MaxHeapL{}

	for key, val := range freq {
		if val > (len(s)+1)/2 {
			return ""
		}
		heap.Push(h, StrSet{key, val})
	}

	var ans []rune

	for h.Len() > 1 {
		ch1, ch2 := heap.Pop(h).(StrSet), heap.Pop(h).(StrSet)
		ans = append(ans, ch1.char, ch2.char)
		ch1.count--
		ch2.count--
		if ch1.count > 0 {
			heap.Push(h, StrSet{ch1.char, ch1.count})
		}
		if ch2.count > 0 {
			heap.Push(h, StrSet{ch2.char, ch2.count})
		}
	}

	if h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(StrSet).char)
	}

	return string(ans)
}

type StrSet struct {
	char  rune
	count int
}

type MaxHeapL []StrSet

func (m MaxHeapL) Len() int           { return len(m) }
func (m MaxHeapL) Less(i, j int) bool { return m[i].count > m[j].count }
func (m MaxHeapL) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *MaxHeapL) Push(x interface{}) {
	*m = append(*m, x.(StrSet))
}
func (m *MaxHeapL) Pop() interface{} {
	old := *m
	n := len(old)
	last := old[n-1]
	*m = old[:n-1]
	return last
}
