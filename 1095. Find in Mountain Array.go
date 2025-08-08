package main

type MountainArray struct {
}

func (this *MountainArray) get(index int) int
func (this *MountainArray) length() int

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := findPeak(mountainArr)
	n := mountainArr.length()

	index := binarySearch(mountainArr, 0, peak, target, true)
	if index != -1 {
		return index
	}

	return binarySearch(mountainArr, peak+1, n-1, target, false)
}

func binarySearch(arr *MountainArray, l, r, target int, ascending bool) int {
	for l <= r {
		mid := (l + r) / 2
		val := arr.get(mid)
		if val == target {
			return mid
		}
		if ascending {
			if val < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if val > target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func findPeak(arr *MountainArray) int {
	l, r := 0, arr.length()-1

	for l < r {
		mid := (l + r) / 2
		if arr.get(mid) < arr.get(mid+1) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l

}
