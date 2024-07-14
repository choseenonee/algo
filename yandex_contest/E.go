package yandex_contest

import (
	"bufio"
	"fmt"
	"os"
)

func E() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &q)

	trie := New()
	poryadokMap := make(map[string]int, n)

	for j := 0; j < n; j++ {
		var key string
		var val int
		fmt.Fscan(in, &key)
		fmt.Fscan(in, &val)
		trie.Add(key, val)
		poryadokMap[key] = j + 1
	}

	var curKey string
	for j := 0; j < q; j++ {
		var flag string
		var key string
		fmt.Fscan(in, &flag)
		if flag == "-" {
			curKey = curKey[:len([]rune(curKey))-1]
		} else {
			fmt.Fscan(in, &key)
		}

		type maxValIdx struct {
			val int
			key string
		}

		curKey = curKey + key
		keyVals := trie.FindKeyValues(curKey)
		ans := maxValIdx{}
		for _, keyVal := range keyVals {
			if keyVal.Value.(int) > ans.val {
				ans.key = keyVal.Key
			}
		}

		ansVal, ok := poryadokMap[ans.key]
		if !ok {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, ansVal)
		}
	}
}
