package main

type empty struct{}

type set map[[2]int]empty

func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights), len(heights[0])
    pac, atl := make(set), make(set)

    var dfs func(r, c int, visit set, prevHeight int)
    dfs = func(r, c int, visit set, prevHeight int) {
        coord := [2]int{r, c}
        _, exists := visit[coord]
        
        if exists || 
        r < 0 || c < 0 || r >= rows || c >= cols ||
        heights[r][c] < prevHeight {
            return
        }

        visit[coord] = empty{}

        dfs(r - 1, c, visit, heights[r][c])
        dfs(r + 1, c, visit, heights[r][c])
        dfs(r, c - 1, visit, heights[r][c])
        dfs(r, c + 1, visit, heights[r][c])
    }

    for r := 0; r < rows; r++ {
        dfs(r, 0, pac, 0)
        dfs(r, cols-1, atl, 0)
    }

    for c := 0; c < cols; c++ {
        dfs(0, c, pac, 0)
        dfs(rows-1, c, atl, 0)
    }

    res := make([][]int, 0)
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            coord := [2]int{r, c}
            _, inPac := pac[coord]
            _, inAtl := atl[coord]

            if inPac && inAtl {
                res = append(res, []int{r, c})
            }
        }
    }

    return res
}