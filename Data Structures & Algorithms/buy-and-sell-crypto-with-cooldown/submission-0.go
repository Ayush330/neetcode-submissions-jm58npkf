func maxProfit(prices []int) int {
    return solve(0, false, true, prices)
}

func solve(day int, isHolding, canBuy bool, prices []int) int{
    if day >= len(prices){
        return 0
    }

    res := solve(day+1, isHolding, true, prices)

    if isHolding{
        res = max(res, prices[day]+solve(day+1, false, false, prices))
    }else if canBuy{
        res = max(res, -prices[day]+solve(day+1, true, true, prices))
    }
    return res
}

func max(a, b int)int{
    if a> b{
        return a
    }
    return b
}