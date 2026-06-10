package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	v        *vertex
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) DecreaseKey(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type vertex struct {
	id,
	index,
	key,
	val int
	e    []edge
	item *Item
}

type edge struct {
	to,
	w int
}

func MSTPrim(l []*vertex) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	total := 0

	v := l[0]
	v.index = -2

	for _, e := range v.e {
		u := l[e.to]
		a := e.w
		u.key = a
		u.val = v.id
		u.item = &Item{v: u, priority: a}
		heap.Push(&pq, u.item)
		u.index = u.item.index
	}

	for pq.Len() > 0 {
		minItem := heap.Pop(&pq).(*Item)
		v = minItem.v

		if v.index == -2 {
			continue
		}

		v.index = -2
		total += minItem.priority

		for _, e := range v.e {
			u := l[e.to]
			a := e.w

			if u.index == -1 {
				u.key = a
				u.val = v.id
				u.item = &Item{v: u, priority: a}
				heap.Push(&pq, u.item)
				u.index = u.item.index
			} else if u.index >= 0 && a < u.key {
				u.key = a
				u.val = v.id
				pq.DecreaseKey(u.item, a)
			}
		}
	}
	return total
}

func main() {
	var (
		n,
		m,
		u,
		v,
		w int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	l := make([]*vertex, n)
	for i := 0; i < n; i++ {
		l[i] = &vertex{index: -1, id: i}
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &u, &v, &w)
		l[u].e = append(l[u].e, edge{to: v, w: w})
		l[v].e = append(l[v].e, edge{to: u, w: w})
	}
	fmt.Println(MSTPrim(l))
}
