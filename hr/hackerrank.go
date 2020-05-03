// Package hackerrank We take sample problem from hackerrank and solve this here.
// We also write tests to check the implementation.
package hackerrank

// CountWord Given a name of a camelCase variable, find out
// the number of words used in the variable.
func CountWord(s string) int {
	var counter int
	for _, char := range s {
		// char below 91 is an uppercase letter
		// given that variable names can't have
		// sepcial character
		if char < 91 {
			counter++
		}
	}
	// +1 because first letter is lowercase
	return counter + 1
}

// CaesarCipher Given a string s of alphabets, and key k (int),
// rotate the ascii value by the number of keys.
// If k tries to go beyond the boundry of alphabet in ascii table,
// rotate it back to the first alphabet e.g. 'y' with
// key 2 will return an 'a'.
func CaesarCipher(input string, k int32) string {
	var final []rune

	// iterate over letters to..
	for _, letter := range input {
		final = append(final, cipher(letter, int(k)))
	}

	return string(final)
}

// cipher func is an intermediate func to initiate rotate func
// based on uppercase or lowercase ascii range. Returns input
// as it is if not found in both ranges.
func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

// rotate takes rune r and calculates the shifted index in ascii table.
func rotate(r rune, base, delta int) rune {
	// 1. Subtracting the base lets up play within 1-26.
	tmp := int(r) - base
	// 2. If we add delta to it. We'll get how many letter to shift.
	// We'll modulo the result by 26, just in case it exceeds 26 limit
	tmp = (tmp + delta) % 26
	// 3. If we add base back to processed result, we get rotated rune
	return rune(tmp + base)
}
