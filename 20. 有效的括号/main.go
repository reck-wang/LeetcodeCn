package main

func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		cur := s[i]
		if cur == '(' || cur == '[' || cur == '{' {
			stack = append(stack, cur)
		} else {
			if len(stack) == 0 {
				return false
			}
			switch cur {
			case ')':
				if stack[len(stack)-1] == '(' {
					return false
				}
			case ']':
				if stack[len(stack)-1] == '[' {
					return false
				}
			case '}':
				if stack[len(stack)-1] == '{' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
