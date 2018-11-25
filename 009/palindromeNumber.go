func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	target := x
	reverse := 0

	for ; x != 0; x /= 10 {
		reverse = (reverse * 10) + (x % 10)
	}

	return reverse == target
}