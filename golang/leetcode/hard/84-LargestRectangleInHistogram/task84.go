package main

func main() {
}

type Height struct {
	height int
	index  int
}

func largestRectangleArea(heights []int) int {
	stack := make([]Height, 0)

	maxArea := 0
	for index, height := range heights {
		last := len(stack) - 1
		if len(stack) == 0 || stack[last].height <= height {
			stack = append(stack, Height{height: height, index: index})
		} else {
			var elem Height

			for len(stack) > 0 && stack[len(stack)-1].height > height {
				elem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				area := (index - elem.index) * elem.height
				if area > maxArea {
					maxArea = area
				}
			}

			stack = append(stack, Height{height: height, index: elem.index})
		}
	}

	for _, height := range stack {
		area := height.height * (len(heights) - height.index)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
