package main

/**
 * @param points: n points on a 2D plane
 * @return: if there is such a line parallel to y-axis that reflect the given points
 */

func IsReflected(points [][]int) bool {
	if points == nil || len(points) == 0 {
		return true
	}
	xMax, xMin := points[0][0], points[0][0]
	for i := 0; i < len(points); i++ {
		if points[i][0] > xMax {
			xMax = points[i][0]
		}
		if points[i][0] < xMin {
			xMin = points[i][0]
		}
	}

	xAvg := float64(xMax+xMin) / 2

	// key = y, value = xAvg - x
	diffs := make(map[int]float64)

	for i := 0; i < len(points); i++ {
		x := points[i][0]
		y := points[i][1]

		diffs[y] += xAvg - float64(x)
	}

	for _, value := range diffs {
		if value != 0 {
			return false
		}
	}

	return true
}
