package main

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	var ans []byte

	var carry byte
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var a byte = '0'
		if i >= 0 {
			a = num1[i]
		}
		var b byte = '0'
		if j >= 0 {
			b = num2[j]
		}
		sum := a + b + carry - '0'
		if sum > '9' {
			carry = 1
			sum -= byte(10)
		} else {
			carry = 0
		}

		ans = append(ans, sum)

	}

	if carry == 1 {
		ans = append(ans, '1')
	}

	for k := 0; k < len(ans)/2; k++ {
		ans[k], ans[len(ans)-k-1] = ans[len(ans)-k-1], ans[k]
	}

	return string(ans)
}
