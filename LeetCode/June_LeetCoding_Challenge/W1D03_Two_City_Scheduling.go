package June_LeetCoding_Challenge

import "sort"

func twoCitySchedCost(costs [][]int) int {
	absFn := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	sort.Slice(costs, func(i, j int) bool {
		return absFn(costs[i][0]-costs[i][1]) > absFn(costs[j][0]-costs[j][1])
	})
	var (
		minCost = 0
		n       = len(costs)
		a       = n / 2
		b       = n / 2
	)
	for i := 0; i < n; i++ {
		if b == 0 || (costs[i][0] <= costs[i][1] && a > 0) {
			a--
			minCost += costs[i][0]
		} else {
			b--
			minCost += costs[i][1]
		}
	}
	return minCost
}
