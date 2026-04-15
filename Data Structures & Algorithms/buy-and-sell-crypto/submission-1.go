func maxProfit(prices []int) int {
	answer := 0
	minm := prices[0]
	for i:=1; i<len(prices); i++{
		minm = min(minm, prices[i])
		answer = max(answer, prices[i]-minm)
	}
	return answer
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}

func min(a, b int)int{
	if a<b{
		return a
	}
	return b
}
