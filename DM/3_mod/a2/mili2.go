package main

import (
	"bufio"
	"fmt"
	"os"
)

type auto struct {
	n, m, q int
	d       [][]int
	f       [][]string
}

func main() {
	var (
		n,
		m,
		q,
		a int
		b string
	)
	alph := "abcdefghijklmnopqrstuvwxyz"
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &q)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a)
			d[i][j] = a
		}
	}
	f := make([][]string, n)
	for i := 0; i < n; i++ {
		f[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &b)
			f[i][j] = b
		}
	}

	au := auto{n: n, m: m, q: q, d: d, f: f}

	fmt.Println("digraph {")
	fmt.Println("    rankdir = LR")

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("    %d -> %d [label = \"%c(%s)\"]\n", i, au.d[i][j],
				rune(alph[j]), au.f[i][j])
		}
	}

	fmt.Println("}")
}
