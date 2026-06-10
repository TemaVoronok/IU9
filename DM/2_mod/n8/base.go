package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type vertex struct {
	x      int
	T1     int
	low    int
	comp   int
	edge   []*vertex
	parent *vertex
}

type condv struct {
	x       int
	visited bool
	edge    []int
	deg     int
}

var (
	time  int = 1
	count int = 1
	s     []*vertex
)

func Find(v *vertex) *vertex {
	if v.parent == v {
		return v
	} else {
		return Find(v.parent)
	}
}

func Union(v, u *vertex) {
	//fmt.Println("объединяем ", v, u)
	rootx := Find(u)
	rooty := Find(v)
	if rootx.x < rooty.x {
		rooty.parent = rootx
	} else {
		rootx.parent = rooty
	}
}

func Tarjan(l []*vertex) {
	s = make([]*vertex, 0)
	for _, v := range l {
		if v.T1 == 0 {
			VisitVertex(l, v)
		}
	}
}

func VisitVertex(l []*vertex, v *vertex) {
	v.T1 = time
	v.low = time
	time++
	s = append(s, v)
	//fmt.Println("вершина", v)
	for _, u := range l[v.x].edge {
		//fmt.Println("ребро к ", u)
		if u.T1 == 0 {
			VisitVertex(l, u)
		}
		if u.comp == 0 && v.low > u.low {
			v.low = u.low
		}
	}
	//fmt.Println("после цикла", v)
	//fmt.Println("стак", s)
	if v.T1 == v.low {
		for {
			u := s[len(s)-1]
			//fmt.Println("работаем с u", u)
			s = s[:len(s)-1]
			u.comp = count
			Union(v, u)
			if u == v {
				break
			}
		}
		count++
	}
}

func main() {
	var (
		n, m, u, v int
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	l := make([]*vertex, n)
	for i := 0; i < n; i++ {
		l[i] = &vertex{x: i, edge: make([]*vertex, 0), T1: 0, comp: 0}
		l[i].parent = l[i]
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &u, &v)
		l[u].edge = append(l[u].edge, l[v])
	}

	Tarjan(l)
	/*fmt.Println("результат")
	for _, v := range l {
		fmt.Println(v.x, Find(v).x, v.low, v.comp)
	}*/

	cond := make(map[int]*condv, 0)
	for _, v := range l {
		_, ok := cond[v.parent.x]
		if !ok {
			cond[v.parent.x] = &condv{x: v.parent.x, deg: 0}
		}
		for _, u := range l[v.x].edge {
			if v.parent != u.parent {
				cond[v.parent.x].edge = append(cond[v.parent.x].edge, u.parent.x)
				_, ok := cond[u.parent.x]
				if !ok {
					cond[u.parent.x] = &condv{x: u.parent.x, deg: 0}
				}
				cond[u.parent.x].deg++
			}
		}
	}
	//fmt.Println("конденсат")
	base := make([]int, 0)
	for _, v := range cond {
		if v.deg == 0 {
			base = append(base, v.x)
		}
	}
	sort.Ints(base)
	for _, i := range base {
		fmt.Print(i, " ")
	}
}

/*
6
6
0 1
1 2
2 0
2 3
4 3
3 5
*/
