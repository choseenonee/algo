package test_contest_ozon

import (
	"bufio"
	"fmt"
	"os"
)

func Stickers() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var defStc string
	var t int

	fmt.Fscan(in, &defStc)
	fmt.Fscan(in, &t)

	runeString := []rune(defStc)

	for i := 0; i < t; i++ {
		var firstIdx, lastIdx int
		var newStc string

		fmt.Fscan(in, &firstIdx, &lastIdx, &newStc)
		newRuneString := []rune(newStc)
		for k := firstIdx; k <= lastIdx; k++ {
			runeString[k-1] = newRuneString[k-firstIdx]
		}
	}

	fmt.Fprint(out, string(runeString))
}
