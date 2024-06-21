package leetcode

func GroupAnagrams(strs []string) [][]string {
	wordsLettersCount := map[[26]byte]int{}
	groupedAnagrams := make([][]string, 0, len(strs))

	for _, word := range strs {
		lettersCount := [26]byte{}
		for _, letter := range word {
			lettersCount[letter-'a']++
		}

		if resIdx, ok := wordsLettersCount[lettersCount]; ok {
			groupedAnagrams[resIdx] = append(groupedAnagrams[resIdx], word)
		} else {
			groupedAnagrams = append(groupedAnagrams, []string{word})
			wordsLettersCount[lettersCount] = len(groupedAnagrams) - 1
		}
	}

	return groupedAnagrams
}
