package June_LeetCoding_Challenge

import "math/rand"

type Solution struct {
	prefixSum []int
}

func Constructor(w []int) Solution {
	var (
		sum       int
		prefixSum = make([]int, len(w)+1)
	)
	for i, e := range w {
		sum += e
		prefixSum[i+1] = sum
	}
	return Solution{prefixSum: prefixSum}
}

func (_this *Solution) PickIndex() int {
	var (
		prefixSumLen = len(_this.prefixSum)
		n            = rand.Intn(_this.prefixSum[prefixSumLen-1])
		l, h         = 0, prefixSumLen - 1
	)
	for l < h-1 {
		mid := (l + h) / 2
		switch {
		case _this.prefixSum[mid] > n:
			h = mid
		case _this.prefixSum[mid] < n:
			l = mid
		default:
			return mid
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
