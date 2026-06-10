package main

import (
	"bufio"
	"fmt"
	"os"
)

type auto struct {
	n, m, q, M int
	d          [][]int
	f          [][]string
	res        map[string]bool
}

type Task struct {
	q int
	s string
}

func (a *auto) BF() {
	e := make([]Task, 0)
	visited := make(map[Task]bool)
	e = append(e, Task{a.q, ""})
	visited[Task{a.q, ""}] = true
	for len(e) != 0 {
		v := e[0]
		e = e[1:]
		if v.s != "" {
			if !a.res[v.s] {
				a.res[v.s] = true
				fmt.Print(v.s, " ")
			}
		}
		if len(v.s) < a.M {
			for i := 0; i < 2; i++ {
				x := a.f[v.q][i]
				u := a.d[v.q][i]
				if x == "-" {
					x = ""
				}
				k := Task{u, v.s + x}
				if !visited[k] {
					visited[k] = true
					e = append(e, Task{u, v.s + x})
				}
			}
		}
	}
}

/*func (a *auto) BFS(q0 int, s string) {
	//fmt.Println("BFS", q0, s)
	if len(s) <= a.M {
		//fmt.Println(s)
		_, ok := a.res[s]
		if !ok && s != "" {
			a.res[s] = true
			fmt.Print(s, " ")
		}
	}
	if len(s) == a.M {
		return
	}
	x := a.f[q0][0]
	u := a.d[q0][0]
	if x == "-" {
		x = ""
	}
	if !(x == "" && q0 == u) {
		a.BFS(u, s+x)
	}
	y := a.f[q0][1]
	v := a.d[q0][1]
	if y == "-" {
		y = ""
	}
	//fmt.Println("второй переход", y, v)
	if !(y == "" && q0 == v) {
		a.BFS(v, s+y)
	}
}*/

func main() {
	var (
		n,
		m,
		q,
		M,
		a int
		b string
	)
	//alph := "abcdefghijklmnopqrstuvwxyz"
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	m = 2
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
	fmt.Fscan(reader, &q, &M)

	au := &auto{n: n, m: m, q: q, d: d, f: f, M: M, res: make(map[string]bool)}

	//fmt.Println(au)

	au.BF()

}
