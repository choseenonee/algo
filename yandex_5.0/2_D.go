package yandex_5_0

import (
	"bufio"
	"fmt"
	"os"
)

func makeBarier(arr [][]int) {
	vals := [2]int{0, 9}

	for _, xv := range vals {
		for k := 0; k < 10; k++ {
			if arr[xv] == nil {
				arr[xv] = make([]int, 10)
			}
			arr[xv][k] = 0
		}
	}
	for _, yv := range vals {
		for k := 0; k < 10; k++ {
			if arr[k] == nil {
				arr[k] = make([]int, 10)
			}
			arr[k][yv] = 0
		}
	}
}

type point struct {
	x int
	y int
}

func startVisiting(visitedMap map[point]struct{}, x, y int, arr [][]int, ans *int) point {
	xs := [4]int{-1, 1, 0, 0}
	ys := [4]int{0, 0, -1, 1}
	var toVisit point
	newSosed := true

	for newSosed {
		var nextX, nextY int
		newSosed = false
		visitedMap[point{x, y}] = struct{}{}
		for j := 0; j < 4; j++ {
			sosed := arr[x+xs[j]][y+ys[j]]
			if sosed == 0 {
				*ans += 1
			}
			if _, visited := visitedMap[point{x + xs[j], y + ys[j]}]; sosed == 1 && !visited {
				if newSosed {
					toVisit = point{nextX, nextY}
				}
				newSosed = true
				nextX = x + xs[j]
				nextY = y + ys[j]
			}
		}
		x = nextX
		y = nextY
	}

	return toVisit
}

func DTwo() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	arr := make([][]int, 10)

	var x, y int
	for l := 0; l < n; l++ {
		fmt.Fscan(in, &x, &y)

		if arr[x] == nil {
			arr[x] = make([]int, 10)
		}

		arr[x][y] = 1
	}

	makeBarier(arr)

	visitedMap := map[point]struct{}{}
	ans := 0

	toVisit := startVisiting(visitedMap, x, y, arr, &ans)
	for toVisit.y != 0 {
		if _, visited := visitedMap[toVisit]; !visited {
			toVisit = startVisiting(visitedMap, toVisit.x, toVisit.y, arr, &ans)
		} else {
			toVisit = point{0, 0}
		}
	}

	fmt.Fprintln(out, ans)
}
