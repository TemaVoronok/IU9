package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type vertex struct {
	v      int
	mark   string
	inc    []int
	parent int
	comp   int
}

type IncidenceList struct {
	inc map[int]*vertex
}

func (l *IncidenceList) AddV(x int) {
	_, ok := l.inc[x]
	if ok {
		panic("vertex already in list")
	}
	l.inc[x] = &vertex{v: x, inc: make([]int, 0)}
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

func DFS1(l *IncidenceList, e *list.List) {
	for _, v := range l.inc {
		v.mark = "white"
	}
	for _, v := range l.inc {
		if v.mark == "white" {
			v.parent = -1
			VisitVertex1(l, v, e)
		}
	}
}

func VisitVertex1(l *IncidenceList, v *vertex, e *list.List) {
	v.mark = "gray"
	e.PushBack(v)
	for _, ed := range v.inc {
		u := l.inc[ed]
		if u.mark == "white" {
			u.parent = v.v
			VisitVertex1(l, u, e)
		}
	}
	v.mark = "black"
}

func DFS2(l *IncidenceList, e *list.List) {
	for _, v := range l.inc {
		v.comp = -1
	}
	comp := 0
	for e.Len() != 0 {
		v := e.Remove(e.Front()).(*vertex)
		if v.comp == -1 {
			VisitVertex2(l, v, comp)
			comp++
		}
	}
}

func VisitVertex2(l *IncidenceList, v *vertex, comp int) {
	v.comp = comp
	for _, ed := range v.inc {
		u := l.inc[ed]
		if u.comp == -1 && u.parent != v.v {
			VisitVertex2(l, u, comp)
		}
	}
}

func CountBridges(l *IncidenceList) int {
	bridgesCount := 0
	for _, v := range l.inc {
		for _, ed := range v.inc {
			u := l.inc[ed]
			if v.comp != u.comp {
				bridgesCount++
			}
		}
	}
	return bridgesCount / 2
}

func main() {
	var (
		n int
		m int
		u int
		v int
		l *IncidenceList = &IncidenceList{make(map[int]*vertex)}
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

	queue := list.New()
	DFS1(l, queue)
	DFS2(l, queue)

	fmt.Println(CountBridges(l))
}
