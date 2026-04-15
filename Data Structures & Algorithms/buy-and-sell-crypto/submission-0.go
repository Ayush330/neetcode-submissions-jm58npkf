func maxProfit(prices []int) int {
	answer := 0
	for i:=0; i<len(prices); i++{
		for j:=i+1; j<len(prices); j++{
			answer = max(answer, prices[j]-prices[i])
		}
	}
	return answer
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
