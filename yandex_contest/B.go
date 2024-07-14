package yandex_contest

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type ByVal []KeyValue

func (a ByVal) Len() int { return len(a) }
func (a ByVal) Less(i, j int) bool {
	return a[i].Value.(int) < a[j].Value.(int)
}
func (a ByVal) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func B() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &q)

	trie := New()

	for j := 0; j < n; j++ {
		var key string
		fmt.Fscan(in, &key)
		trie.Add(key, j+1)
	}

	for j := 0; j < q; j++ {
		var number int
		var key string
		fmt.Fscan(in, &number)
		fmt.Fscan(in, &key)

		keyVals := trie.FindKeyValues(key)

		sort.Sort(ByVal(keyVals))

		fmt.Fprintln(out, keyVals)

		if number-1 >= len(keyVals) {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, keyVals[number-1].Value.(int))
		}
	}
}
