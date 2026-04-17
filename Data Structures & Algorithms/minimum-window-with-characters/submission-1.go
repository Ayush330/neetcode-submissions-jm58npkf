func minWindow(s string, t string) string {
    refrenceCount := len(t)
	left := 0
	//right := 0
	mp := make(map[byte]int)
	for i:=0; i<len(t); i++{
		mp[t[i]]++
	}

	runningMp := make(map[byte]int)

	runningLength := refrenceCount
	minmLength := len(s)+1
	start, end := -1, -1
	var right int
	for right = 0; right < len(s); right++{
		char := s[right]
		if target, ok := mp[char]; ok {//&& runningMp[char] < target{
			runningMp[char]++
			if runningMp[char] <= target{
				runningLength--
			}
		}
		//fmt.Println("1: The running length is: ", runningLength, " and the length of the actual string is ", len(s))
		// check if valid 
		if runningLength == 0{
			if right-left+1 < minmLength{
				minmLength = right-left+1
				start = left
				end = right
			}
			for runningLength == 0{
				charInTalks := s[left]
				left++
				if target, ok := mp[charInTalks]; ok{
					runningMp[charInTalks]--
					if target > runningMp[charInTalks]{
						runningLength++
					}else{
						if right-left+1 < minmLength{
						minmLength = right-left+1
						start = left
						end = right
						}
					}
				}else{
					if right-left+1 < minmLength{
						minmLength = right-left+1
						start = left
						end = right
					}
				}
				//fmt.Println("2: The running length is: ", runningLength)
			}
		}
	}
	//fmt.Println("Start Is: ", start, " and End is: ", end)
	if runningLength == 0{
		if right-left+1 < minmLength{
				minmLength = right-left+1
				start = left
				end = right
			}
	}
	if start == -1{
		return ""
	}
	return s[start:end+1]
}


