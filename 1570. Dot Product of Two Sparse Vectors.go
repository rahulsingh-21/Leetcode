package main

type SparseVector struct {
	m map[int]int
}

func ConstructorV(nums []int) SparseVector {
	m := make(map[int]int)

	for i, val := range nums {
		if val != 0 {
			m[i] = val
		}
	}
	return SparseVector{m}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {

	dotProduct := 0
	for key, val := range this.m {
		if n, ok := vec.m[key]; ok {
			dotProduct += val * n
		}
	}
	return dotProduct
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
