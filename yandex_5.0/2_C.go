package yandex_5_0

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func CTwo() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, 0, n)

	for k := 0; k < n; k++ {
		var elem int
		fmt.Fscan(in, &elem)
		arr = append(arr, elem)
	}

	verevka := slices.Max(arr)

	var partsSum int
	var happened = false

	for _, val := range arr {
		if val == verevka && !happened {
			happened = true
			continue
		}
		if val != verevka || happened {
			partsSum += val
		}
	}

	if verevka == partsSum || partsSum > verevka {
		fmt.Fprint(out, verevka+partsSum)
	} else {
		fmt.Fprint(out, verevka-partsSum)
	}
}
