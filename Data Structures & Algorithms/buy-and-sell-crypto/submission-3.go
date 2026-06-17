func maxProfit(prices []int) int {
	answer := 0
	runningMin := prices[0]
	for i:=1; i<len(prices); i++{
		answer = max(answer, prices[i]-runningMin)
		runningMin = min(runningMin, prices[i])
	}
	return answer
}

func min(a, b int)int{
	if a<b{
		return a
	}
	return b
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}