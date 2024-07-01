package yandex_5_0

import (
	"bufio"
	"fmt"
	"os"
)

// учесть что K не гарантированно меньше N

func BTwo() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var k int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)

	arr := make([]int, 0, n)

	for d := 0; d < n; d++ {
		var elem int
		fmt.Fscan(in, &elem)
		arr = append(arr, elem)
	}

	var ans int

	for i := 0; i < n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			a := arr[i]
			b := arr[i-j]
			_ = a
			_ = b
			ans = max(ans, arr[i]-arr[i-j])
		}
	}

	fmt.Fprint(out, ans)
}
