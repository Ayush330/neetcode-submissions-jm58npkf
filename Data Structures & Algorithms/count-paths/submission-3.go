type point struct{
    a int
    b int
}

func uniquePaths(m int, n int) int {
    memo := make(map[point]int)
    return solve(m, n, 0, 0, memo)
    //return memo[point{m-1, n-1}]
}


func solve(m, n, x, y int, memo map[point]int) int{

    if curr, ok := memo[point{x,y}]; ok != false{
        return curr
    }

    if x >= m || y >= n{
        return 0
    }

    if x == m-1 && y == n-1{
        return 1
    }

    memo[point{x, y}] = solve(m, n, x+1, y, memo) + solve(m, n, x, y+1, memo)
    return memo[point{x, y}]
}
