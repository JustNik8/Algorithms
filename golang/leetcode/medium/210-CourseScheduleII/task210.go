package main

type empty struct{}

type set map[int]empty

func findOrder(numCourses int, prerequisites [][]int) []int {
	size := len(prerequisites)
	pres := make(map[int][]int)

	ans := make([]int, 0)

	for i := 0; i < size; i++ {
		parent, child := prerequisites[i][0], prerequisites[i][1]

		_, exists := pres[parent]
		if !exists {
			pres[parent] = make([]int, 0)
		}

		pres[parent] = append(pres[parent], child)
	}

	visit, cycle := make(set), make(set)

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if _, inCycle := cycle[crs]; inCycle {
			return false
		}
		if _, visited := visit[crs]; visited {
			return true
		}

		cycle[crs] = empty{}
		visit[crs] = empty{}

		for _, pre := range pres[crs] {
			ok := dfs(pre)
			if !ok {
				return false
			}
		}

		delete(cycle, crs)
		ans = append(ans, crs)

		return true
	}

	for crs := 0; crs < numCourses; crs++ {
		ok := dfs(crs)
		if !ok {
			return []int{}
		}
	}

	return ans
}
