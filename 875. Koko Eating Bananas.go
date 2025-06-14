package main

func minEatingSpeed(piles []int, h int) int {

	maxi := 0
	for _, pile := range piles {
		maxi = max(maxi, pile)
	}

	l := 1
	r := maxi

	for l < r {
		mid := (l + r) / 2
		ho := 0
		for _, pile := range piles {
			if pile%mid == 0 {
				ho += pile / mid
			} else {
				ho += pile/mid + 1
			}
		}

		if ho <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
