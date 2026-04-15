func isValid(s string) bool {
    stack := []byte{}
	for i:=0; i<len(s); i++{
		if s[i] == '(' || s[i] == '{' || s[i] == '['{
			stack = append(stack, s[i])
		}else{
			lenStack := len(stack)
			if lenStack > 0{
				topEl := stack[lenStack-1]
				if s[i] == ')' && topEl == '(' || s[i] == '}' && topEl == '{' || s[i] == ']' && topEl == '['{
					stack = stack[:lenStack-1]
				}else{
					// return false
					stack = append(stack, s[i])
				}
			}else{
				stack = append(stack, s[i])
			}
		}
	}
	return len(stack) == 0
}
