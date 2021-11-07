package palindrome

import (
	"fmt"
	"unicode/utf8"
)

func reverse(rev string) string {
	runes := []rune(rev)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Palindrome(param ...string) {
	for _, str := range param {
		if str == "" {
			continue
		}
		n := utf8.RuneCountInString(str)
		rev := ""
		for i := range str {
			rev = reverse(str[i:])
			if rev == str[i:] {
				rev = reverse(str[:i])
				n += utf8.RuneCountInString(rev)
				break
			}
		}
		fmt.Println(str+rev, n)
	}
}
