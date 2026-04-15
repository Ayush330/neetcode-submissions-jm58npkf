func trap(height []int) int {
	l := 1
	r := len(height)-2
	lm := height[l-1]
	rm := height[r+1]
	answer := 0
	for l<=r{
		if lm < rm{
			answer += max(0, lm-height[l])
			lm = max(lm, height[l])
			l++
		}else{
			answer += max(0, rm-height[r])
			rm = max(rm, height[r])
			r--
		}
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
