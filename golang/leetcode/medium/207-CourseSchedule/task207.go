package main

type empty struct{}
type set map[int]empty

func canFinish(numCourses int, prerequisites [][]int) bool {
	size := len(prerequisites)
	pres := make(map[int][]int)

	for i := 0; i < size; i++ {
		parent, child := prerequisites[i][0], prerequisites[i][1]

		_, exists := pres[parent]
		if !exists {
			pres[parent] = make([]int, 0)
		}

		pres[parent] = append(pres[parent], child)
	}

	visit := make(set)

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if _, visited := visit[crs]; visited {
			return false
		}
		if len(pres[crs]) == 0 {
			return true
		}

		visit[crs] = empty{}
		for _, pre := range pres[crs] {
			ok := dfs(pre)
			if !ok {
				return false
			}
		}

		pres[crs] = []int{}
		delete(visit, crs)

		return true
	}

	for crs := 0; crs < numCourses; crs++ {
		ok := dfs(crs)
		if !ok {
			return false
		}
	}

	return true
}
