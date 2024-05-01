package test_contest_ozon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func isEqual(currentTiming, lastTiming int) bool {
	return currentTiming-1 == lastTiming || currentTiming == lastTiming || currentTiming+1 == lastTiming
}

func Championship() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tries int
	fmt.Fscan(in, &tries)
	for d := 0; d < tries; d++ {

		var runnersCount int
		fmt.Fscan(in, &runnersCount)

		runnersTiming := make([]int, runnersCount)
		for i := 0; i < runnersCount; i++ {
			var timing int
			fmt.Fscan(in, &timing)
			runnersTiming[i] = timing // O(n)
		}

		sortedRunnersTiming := make([]int, runnersCount)
		copy(sortedRunnersTiming, runnersTiming)
		sort.Ints(sortedRunnersTiming) // O(n log n)

		timingPlaceMap := map[int]int{}

		var lastPlace, lastTiming, skippedPlaces int
		lastPlace = 1
		skippedPlaces = 0
		lastTiming = sortedRunnersTiming[0]
		timingPlaceMap[sortedRunnersTiming[0]] = 1
		for k := 1; k < len(sortedRunnersTiming); k++ {
			currentTiming := sortedRunnersTiming[k]
			if isEqual(currentTiming, lastTiming) {
				skippedPlaces++
				lastTiming = currentTiming
				timingPlaceMap[currentTiming] = lastPlace
			} else {
				lastPlace++
				lastTiming = currentTiming
				timingPlaceMap[currentTiming] = lastPlace + skippedPlaces
				lastPlace += skippedPlaces
				skippedPlaces = 0
			}
		}

		for _, i := range runnersTiming {
			fmt.Fprint(out, timingPlaceMap[i], " ")
		}
		fmt.Fprintln(out, "")
	}
}
