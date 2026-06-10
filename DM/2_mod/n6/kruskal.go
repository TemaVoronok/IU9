package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Task struct {
	start int
	end   int
}

type Stack struct {
	st  []Task
	cnt int
}

func (s *Stack) Init() {
	s = &Stack{make([]Task, 0), 0}
}

func (s *Stack) IsEmpty() bool {
	return s.cnt == 0
}

func (s *Stack) Push(x Task) {
	s.st = append(s.st, x)
	s.cnt++
}

func (s *Stack) Pop() Task {
	if s.cnt != 0 {
		s.cnt--
		item := s.st[s.cnt]
		s.st = s.st[:s.cnt]
		return item
	}
	return Task{}
}

func qssort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	s := Stack{}
	s.Init()
	s.Push(Task{0, n - 1})
	for !s.IsEmpty() {
		curTask := s.Pop()
		low := curTask.start
		i := low
		high := curTask.end
		if low >= high {
			continue
		}
		for i < high {
			if less(i, high) {
				swap(i, low)
				low++
			}
			i++
		}
		swap(low, high)
		s.Push(Task{curTask.start, low - 1})
		s.Push(Task{low + 1, high})
	}
}

type vertex struct {
	v      int
	x      float64
	y      float64
	parent int
	countV int
}

type edge struct {
	v, u int
	dist float64
}

type Sets struct {
	y []*vertex
}

func (s *Sets) Find(x int) int {
	i := s.y[x]
	if i.parent == x {
		return x
	}
	root := s.Find(i.parent)
	i.parent = root
	return root
}

func (s *Sets) Union(x, y int) {
	i := s.Find(x)
	j := s.Find(y)
	k := s.y[i]
	f := s.y[j]
	if s.Find(x) == s.Find(y) {
		return
	}
	if k.countV < f.countV {
		k, f = f, k
		i, j = j, i
	}
	k.parent = j
}

func (s *Sets) SpanningTree(e []edge) []edge {
	mst := make([]edge, 0)
	for i := 0; i < len(e); i++ {
		edge := e[i]
		if s.Find(edge.v) != s.Find(edge.u) {
			s.Union(edge.v, edge.u)
			mst = append(mst, edge)
		}
	}
	return mst
}

func main() {
	var (
		n    int
		x, y float64
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	var g Sets = Sets{make([]*vertex, n)}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x, &y)
		g.y[i] = &vertex{v: i, x: x, y: y, parent: i, countV: 1}
	}
	var e []edge = make([]edge, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := math.Sqrt(math.Pow(g.y[i].x-g.y[j].x, 2) + math.Pow(g.y[i].y-g.y[j].y, 2))
			e = append(e, edge{v: i, u: j, dist: d})
		}
	}

	less := func(i, j int) bool {
		return e[i].dist < e[j].dist
	}
	swap := func(i, j int) {
		e[i], e[j] = e[j], e[i]
	}

	qssort(len(e), less, swap)

	s := 0.0
	for _, i := range g.SpanningTree(e) {
		s += i.dist
	}
	fmt.Printf("%.2f", s)

}
