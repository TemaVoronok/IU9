package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to *vertex
	c  int
}

type vertex struct {
	v    int
	edge []*edge
	way  []int
}

func Lex(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return false
}

func BFS(l []*vertex) {
	q := make([]*vertex, 0)
	w := l[0]
	q = append(q, w)
	//i := 0
	//fmt.Println("BFS")
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		//fmt.Println("Сейчас ", v)
		for _, k := range l[v.v].edge {
			u := k.to

			newWay := make([]int, len(v.way)+1)
			copy(newWay, v.way)
			newWay[len(v.way)] = k.c
			if len(u.way) == 0 ||
				(len(newWay) == len(u.way) && Lex(newWay, u.way)) ||
				len(newWay) < len(u.way) {
				u.way = newWay
				q = append(q, u)
			}
		}
	}
}

func main() {
	var (
		n, m, a, b, c int
		l             []*vertex
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	l = make([]*vertex, n)
	for i := 0; i < n; i++ {
		l[i] = &vertex{v: i, edge: make([]*edge, 0), way: make([]int, 0)}
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b, &c)
		a--
		b--
		l[a].edge = append(l[a].edge, &edge{to: l[b], c: c})
		l[b].edge = append(l[b].edge, &edge{to: l[a], c: c})
	}

	BFS(l)

	/*fmt.Println("резульат")
	for i, u := range l {
		fmt.Println("Вершина ", i, u)
	}*/
	res := l[n-1].way
	fmt.Println(len(res))
	for _, i := range res {
		fmt.Print(i, " ")
	}
}
