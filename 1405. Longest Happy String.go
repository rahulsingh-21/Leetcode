package main

import "container/heap"

func longestDiverseString(a int, b int, c int) string {

	h := &maxHeapII{}
	if a > 0 {
		heap.Push(h, charSet{'a', a})
	}
	if b > 0 {
		heap.Push(h, charSet{'b', b})
	}
	if c > 0 {
		heap.Push(h, charSet{'c', c})
	}

	var result []byte
	for h.Len() > 0 {
		first := heap.Pop(h).(charSet)

		if len(result) >= 2 && first.char == result[len(result)-1] && first.char == result[len(result)-2] {
			if h.Len() == 0 {
				break
			}
			second := heap.Pop(h).(charSet)
			result = append(result, second.char)
			second.count--
			if second.count > 0 {
				heap.Push(h, second)
			}
			heap.Push(h, first)
		} else {
			result = append(result, first.char)
			first.count--
			if first.count > 0 {
				heap.Push(h, first)
			}
		}
	}
	return string(result)
}

type charSet struct {
	char  byte
	count int
}

type maxHeapII []charSet

func (m maxHeapII) Len() int           { return len(m) }
func (m maxHeapII) Less(i, j int) bool { return m[i].count > m[j].count }
func (m maxHeapII) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeapII) Push(x interface{}) {
	*m = append(*m, x.(charSet))
}
func (m *maxHeapII) Pop() interface{} {
	old := *m
	last := old[len(old)-1]
	*m = old[:len(old)-1]
	return last
}
