package yandex_5_0

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type berry struct {
	up     int
	down   int
	number int
}

type ByDiff []berry

func (a ByDiff) Len() int { return len(a) }
func (a ByDiff) Less(i, j int) bool {
	return a[i].up-a[i].down > a[j].up-a[j].down
}
func (a ByDiff) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByUp []berry

func (a ByUp) Len() int { return len(a) }
func (a ByUp) Less(i, j int) bool {
	if a[i].up-a[i].down < 0 && a[j].up-a[j].down < 0 {
		return a[i].up > a[j].up
	}
	return a[i].up-a[i].down > a[j].up-a[j].down
}
func (a ByUp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func ETwo() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	arr := make([]berry, 0, n)

	for k := 0; k < n; k++ {
		var up, down int
		fmt.Fscan(in, &up, &down)
		arr = append(arr, berry{up, down, k + 1})
	}

	sort.Sort(ByDiff(arr))
	sort.Sort(ByUp(arr))

	fmt.Fprintln(out, arr)

	var ans int
	var curHeight int
	for _, val := range arr {
		curHeight += val.up
		ans = max(ans, curHeight)
		curHeight -= val.down
	}

	fmt.Fprintln(out, ans)
	for _, val := range arr {
		fmt.Fprint(out, val.number, " ")
	}
}
