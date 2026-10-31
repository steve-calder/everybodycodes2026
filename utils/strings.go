package utils

func ReverseString(s string) string {
	// 1. Convert string to a slice of runes to safely handle multi-byte characters
	runes := []rune(s)

	// 2. Reverse the slice in-place using a two-pointer approach
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 3. Convert the rune slice back into a valid Go string
	return string(runes)
}
