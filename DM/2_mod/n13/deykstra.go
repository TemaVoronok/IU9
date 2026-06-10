package main

import (
	"fmt"
	"math"

	//"math"
	"bufio"
	"os"
)

type vertex struct {
	dist   int
	parent *vertex
	w      int
	index  int
	x      int
	y      int
}

func Heapify(i int, p []*vertex) {
	for {
		l := 2*i + 1
		r := l + 1
		j := i
		if l < len(p) && p[i].dist > p[l].dist {
			i = l
		}
		if r < len(p) && p[i].dist > p[r].dist {
			i = r
		}
		if i == j {
			break
		}
		p[i], p[j] = p[j], p[i]
		p[i].index = i
		p[j].index = j
	}
}

type PriorityQueue struct {
	heap  []*vertex
	cap   int
	count int
}

func (q *PriorityQueue) Minimum() *vertex {
	return q.heap[0]
}

func (q *PriorityQueue) Empty() bool {
	return q.count == 0
}

func (q *PriorityQueue) Insert(v *vertex) {
	i := q.count
	q.count++
	q.heap[i] = v
	for i > 0 && q.heap[(i-1)/2].dist > q.heap[i].dist {
		q.heap[(i-1)/2], q.heap[i] = q.heap[i], q.heap[(i-1)/2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

func (q *PriorityQueue) ExtractMin() *vertex {
	v := q.heap[0]
	q.count--
	if q.count > 0 {
		q.heap[0] = q.heap[q.count]
		q.heap[0].index = 0
		Heapify(0, q.heap[:q.count])
	}
	return v
}

func (q *PriorityQueue) DecreaseKey(u *vertex) {
	i := u.index
	for i > 0 && q.heap[(i-1)/2].dist > u.dist {
		q.heap[(i-1)/2], q.heap[i] = q.heap[i], q.heap[(i-1)/2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

func Relax(u, v *vertex, w int) bool {
	changed := u.dist+w < v.dist
	if changed {
		v.dist = w + u.dist
		v.parent = u
	}
	return changed
}

func Dijkstra(l []*vertex) {
	n := len(l)
	q := &PriorityQueue{heap: make([]*vertex, n), cap: n, count: 0}
	l[0].dist = l[0].w
	for _, i := range l {
		q.Insert(i)
	}
	for !q.Empty() {
		v := q.ExtractMin()
		//fmt.Println(v)
		v.index = -1
		m := int(math.Sqrt(float64(n)))
		//fmt.Println("neib")
		if v.y+1 < m {
			u := l[v.x*m+(v.y+1)]
			/*fmt.Println("было", u)
			for i := 0; i < 5; i++ {
				fmt.Println(q.heap[i])
			}*/
			if u.index != -1 && Relax(v, u, u.w) {
				q.DecreaseKey(u)
			}
			/*fmt.Println("стало", u)
			for i := 0; i < 5; i++ {
				fmt.Println(q.heap[i])
			}*/
		}
		if v.x+1 < m {
			u := l[(v.x+1)*m+v.y]
			/*fmt.Println("было", u)
			for i := 0; i < 5; i++ {
				fmt.Println(q.heap[i])
			}*/
			if u.index != -1 && Relax(v, u, u.w) {
				q.DecreaseKey(u)
			}
			/*fmt.Println("стало", u)
			for i := 0; i < 5; i++ {
				fmt.Println(q.heap[i])
			}*/
		}
		if v.x-1 >= 0 {
			u := l[(v.x+-1)*m+v.y]
			if u.index != -1 && Relax(v, u, u.w) {
				q.DecreaseKey(u)
			}
		}
		if v.y-1 >= 0 {
			u := l[v.x*m+(v.y-1)]
			if u.index != -1 && Relax(v, u, u.w) {
				q.DecreaseKey(u)
			}
		}
		//fmt.Println("Очередь после ", v)
		/*
			for i := 0; i < 25; i++ {
				fmt.Println(q.heap[i])
			}*/
		//fmt.Println("next")
	}
}

func main() {
	var (
		n int
		w int
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	l := make([]*vertex, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &w)
			l[i*n+j] = &vertex{dist: 999999, parent: nil, w: w, x: i, y: j}
		}
	}
	Dijkstra(l)
	fmt.Println(l[n*n-1].dist)
}
