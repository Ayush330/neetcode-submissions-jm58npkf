func largestRectangleArea(h []int) int {
	l := len(h)
	var ans int
	for i:=0; i<l; i++{
		left, right := i, i
		// calculate left 
		for j:=i; j >= 0; j--{
			if h[j] < h[i]{
				break
			}
			left = j
		}

		//calculate right

		for j:=i; j < l; j++{
			if h[j] < h[i]{
				break
			}
			right = j
		}
		ans = max((right-left+1) * h[i], ans)
		//fmt.Printf("CurrentIndex: %d\tLeft: %d\tRight: %d\tAnswer:%d\n", i, left, right, ans)
	}

	return ans
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
