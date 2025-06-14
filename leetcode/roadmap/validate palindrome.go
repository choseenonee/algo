package leetcode

import (
	"strings"
)

// 1040 1103
// 48 57
// 97 122
// 65 90

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	processedRunes := make([]rune, 0, len(s))
	for _, val := range s {
		if val >= 1040 && val <= 1103 || val >= 48 && val <= 57 || val >= 97 && val <= 122 || val >= 65 && val <= 90 {
			processedRunes = append(processedRunes, val)
		}
	}

	for pointer := 0; pointer < len(processedRunes)/2; pointer++ {
		if processedRunes[pointer] != processedRunes[len(processedRunes)-pointer-1] {
			return false
		}
	}

	return true
}
