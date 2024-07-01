package yandex_5_0

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ATwo() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int
	fmt.Fscan(in, &k)

	var maxY = -int(math.Pow10(10))
	var minY = int(math.Pow10(10))

	var maxX = -int(math.Pow10(10))
	var minX = int(math.Pow10(10))

	for ; k > 0; k-- {
		var x, y int
		fmt.Fscan(in, &x, &y)
		maxY = max(maxY, y)
		minY = min(minY, y)
		maxX = max(maxX, x)
		minX = min(minX, x)
	}

	fmt.Println(minX, minY, maxX, maxY)
}
