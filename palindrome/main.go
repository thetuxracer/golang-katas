package main

import "fmt"

type Word struct {
	word string
	isPalindrome bool
}

func NewPalindrome(p string, is bool) *Word {
	left := 0
	right := len(p) - 1

	for left < right {
		if p[left] == p[right] {
			is = true
		}
		left++
		right--
	}
	return &Word{word: p, isPalindrome: is}
}

func (p Word) formatOutput() string {
	if p.isPalindrome {
		return fmt.Sprintf("%s is a palindrome.", p.word)
	} else {
		return fmt.Sprintf("%s is not a palindrome.", p.word)
	}
}

func main() {
	fmt.Println(NewPalindrome("racecar", false).formatOutput())
	fmt.Println(NewPalindrome("noon", false).formatOutput())
	fmt.Println(NewPalindrome("hello", false).formatOutput())
	fmt.Println(NewPalindrome("----", false).formatOutput())
}