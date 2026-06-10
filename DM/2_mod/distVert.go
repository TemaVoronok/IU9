package main

import (
	"bufio"
	"fmt"
	"os"
)

type vertex struct {
	v   int
	inc []int
}

type IncidenceList struct {
	inc       map[int]*vertex
	oporSet   map[int]bool
	oporCount int
}

func (l *IncidenceList) AddV(x int) {
	_, ok := l.inc[x]
	if ok {
		panic("vertex already in list")
	}
	l.inc[x] = &vertex{v: x, inc: make([]int, 0)}
	l.oporSet[x] = false
}

func (l *IncidenceList) AddE(u, v int) {
	_, ok1 := l.inc[u]
	_, ok2 := l.inc[v]
	if !ok1 || !ok2 {
		panic("vertex not in list")
	}
	l.inc[u].inc = append(l.inc[u].inc, v)
	l.inc[v].inc = append(l.inc[v].inc, u)
}

func (l *IncidenceList) BFS(n int) {
	distFrom := make([][]int, 0, l.oporCount)

	queue := make([]int, 0, n)

	for opor := range l.oporSet {
		if !l.oporSet[opor] {
			continue
		}
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[opor] = 0
		queue = queue[:0]
		queue = append(queue, opor)
		head := 0
		for head < len(queue) {
			v := queue[head]
			head++
			for _, uv := range l.inc[v].inc {
				if dist[uv] == -1 {
					dist[uv] = dist[v] + 1
					queue = append(queue, uv)
				}
			}
		}
		distFrom = append(distFrom, dist)
	}

	flag := true
	for v := 0; v < n; v++ {
		if l.oporSet[v] {
			continue
		}
		d0 := distFrom[0][v]
		if d0 == -1 {
			continue
		}
		ok := true
		for i := 1; i < len(distFrom); i++ {
			if distFrom[i][v] != d0 {
				ok = false
				break
			}
		}
		if ok {
			fmt.Print(v, " ")
			flag = false
		}
	}
	if flag {
		fmt.Println("-")
	}
}
func main() {
	var (
		n int
		m int
		u int
		v int
		k int
		l *IncidenceList = &IncidenceList{inc: make(map[int]*vertex), oporSet: make(map[int]bool)}
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)
	for i := 0; i < n; i++ {
		l.AddV(i)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &u, &v)
		l.AddE(u, v)
	}
	fmt.Fscan(reader, &k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &u)
		l.oporCount++
		l.oporSet[u] = true
	}

	/*for j, i := range l.inc {
		fmt.Println(j, i)
	}
	fmt.Println(opor)*/
	l.BFS(n)
}
