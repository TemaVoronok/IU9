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

func (a auto) DFS() map[int]int {
	t := make(map[int]int)
	a.VisitVertex(a.q, t)
	return t
}

func (a auto) VisitVertex(v int, t map[int]int) {
	t[v] = len(t)
	for i := 0; i < a.m; i++ {
		u := a.d[v][i]
		if _, ok := t[u]; ok == false {
			a.VisitVertex(u, t)
		}
	}
}

func main() {
	var (
		n,
		m,
		q,
		a int
		b string
	)
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
	t := au.DFS()

	newA := auto{n: au.n, m: au.m, q: t[au.q]}
	newD := make([][]int, n)
	newF := make([][]string, n)

	for i := 0; i < n; i++ {
		newD[i] = make([]int, m)
		newF[i] = make([]string, m)
	}

	for i := 0; i < n; i++ {
		v, _ := t[i]
		newF[v] = f[i]
		for j := 0; j < m; j++ {
			newD[v][j] = t[d[i][j]]
		}
	}

	newA.d = newD
	newA.f = newF

	fmt.Println(newA.n)
	fmt.Println(newA.m)
	fmt.Println(newA.q)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(newA.d[i][j], " ")
		}
		fmt.Print("\n")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(newA.f[i][j], " ")
		}
		fmt.Print("\n")
	}

}
