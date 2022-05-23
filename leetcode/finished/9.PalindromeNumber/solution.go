package solution

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	maxMultiple := getMaxMultiple(x)
	for {
		if x < 10 && maxMultiple == 1 {
			return true
		}
		if x/maxMultiple != x%10 {
			return false
		}
		x = x - maxMultiple*(x/maxMultiple)
		x = x / 10
		if x == 0 {
			return true
		}
		maxMultiple = maxMultiple / 100
	}
}

func getMaxMultiple(x int) int {
	v := 1
	for {
		if x < 10 {
			break
		}

		if x/10 != 0 {
			v *= 10
		}
		x = x / 10
	}
	return v
}
