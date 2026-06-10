package main

import (
	"bufio"
	"fmt"
	"os"
)

type auto struct {
	q     []*pos
	m, q0 int
	d     [][]int
	f     [][]string
}

type pos struct {
	q      int
	parent *pos
	depth  int
}

func (b *pos) Find() *pos {
	if b.parent == b {
		return b
	}
	return b.parent.Find()
}

func (b *pos) Union(c *pos) {
	rootb := b.Find()
	rootc := c.Find()
	rootb.parent = rootc
}

func (a *auto) Split1() ([]*pos, int) {
	m := len(a.q)
	for i := 0; i < m; i++ {
		a.q[i] = &pos{depth: 0, q: i}
		a.q[i].parent = a.q[i]
	}
	pi := make([]*pos, m)
	for i := 0; i < len(a.q); i++ {
		for j := i + 1; j < len(a.q); j++ {
			q1, q2 := a.q[i], a.q[j]
			if q1.Find() != q2.Find() {
				eq := true
				for k := 0; k < a.m; k++ {
					if a.f[q1.q][k] != a.f[q2.q][k] {
						eq = false
						break
					}
				}
				if eq {
					q1.Union(q2)
					m--
				}
			}
		}
	}
	for _, q := range a.q {
		pi[q.q] = q.Find()
	}
	return pi, m
}

func (a *auto) Split(pi []*pos) ([]*pos, int) {
	m := len(a.q)
	for i := 0; i < m; i++ {
		a.q[i] = &pos{depth: 0, q: i}
		a.q[i].parent = a.q[i]
	}
	for i := 0; i < len(a.q); i++ {
		for j := i + 1; j < len(a.q); j++ {
			q1, q2 := a.q[i], a.q[j]
			if pi[q1.q] == pi[q2.q] && q1.Find() != q2.Find() {
				eq := true
				for k := 0; k < a.m; k++ {
					w1, w2 := a.d[q1.q][k], a.d[q2.q][k]
					if pi[w1] != pi[w2] {
						eq = false
						break
					}
				}
				if eq {
					q1.Union(q2)
					m--
				}
			}
		}
	}
	for _, q := range a.q {
		pi[q.q] = q.Find()
	}
	return pi, m
}

func (a *auto) AufenkampHohn() auto {
	pi, m := a.Split1()
	var m1 int
	for {
		pi, m1 = a.Split(pi)
		if m1 == m {
			break
		}
		m = m1
	}

	classIndex := make(map[*pos]int)
	mapping := make(map[int]int)
	Q1 := make([]*pos, m)
	idx := 0

	var dfs func(currOld int)
	dfs = func(currOld int) {
		rep := pi[currOld]
		if _, ok := classIndex[rep]; ok {
			return
		}
		currNew := idx
		classIndex[rep] = currNew
		mapping[currNew] = currOld
		Q1[currNew] = rep
		idx++
		for i := 0; i < a.m; i++ {
			dfs(a.d[currOld][i])
		}
	}

	dfs(a.q0)

	d1 := make([][]int, m)
	f1 := make([][]string, m)
	for j := 0; j < m; j++ {
		d1[j] = make([]int, a.m)
		f1[j] = make([]string, a.m)
	}

	for i := 0; i < m; i++ {
		oldState := mapping[i]
		for k := 0; k < a.m; k++ {
			toOld := a.d[oldState][k]
			toRep := pi[toOld]
			d1[i][k] = classIndex[toRep]
			f1[i][k] = a.f[oldState][k]
		}
	}

	newq0 := classIndex[pi[a.q0]]
	return auto{q: Q1, d: d1, f: f1, q0: newq0}
}

func main() {
	var (
		n, m, q, a int
		b          string
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

	au := auto{q: make([]*pos, n), m: m, q0: q, d: d, f: f}

	e := au.AufenkampHohn()
	fmt.Println("digraph {")
	fmt.Println("    rankdir = LR")

	for i := 0; i < len(e.q); i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("    %d -> %d [label = \"%c(%s)\"]\n", i, e.d[i][j],
				rune(alph[j]), e.f[i][j])
		}
	}

	fmt.Println("}")
}
