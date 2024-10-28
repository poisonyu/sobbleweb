package word

import "unicode"

func IsPalindrome(s string) bool {
	// var letters []rune = []rune(s)
	// var letters []rune
	// 给letters预分配一个容量足够大的数组
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	// for i := range letters {
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
