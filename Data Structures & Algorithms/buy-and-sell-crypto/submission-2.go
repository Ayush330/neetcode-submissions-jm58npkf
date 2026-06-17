func maxProfit(prices []int) int {
	answer := 0
	movingMax := make([]int, len(prices))
	movingMax[len(prices)-1] = prices[len(prices)-1]
	for i:=len(prices)-2; i>=0; i--{
		movingMax[i] = max(movingMax[i+1], prices[i])
	}

	for i:=0; i<len(prices)-1; i++{
		answer = max(answer, movingMax[i]-prices[i])
	}

	return answer
}


func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}