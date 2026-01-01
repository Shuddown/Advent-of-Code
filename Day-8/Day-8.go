package main

import (
	"bufio"
	"cmp"
	"container/heap"
	"fmt"
	"github.com/Shuddown/Advent-of-Code/utils"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

const (
	MAX_HEAP_SIZE = 10000
)

type point struct {
	x int
	y int
	z int
}

type Pair struct {
	p1, p2 *point
	d      float64
}

type DSU[T comparable] struct {
	parent  map[T]T
	size    map[T]int
	maxSize int
}

func NewDSU[T comparable]() *DSU[T] {
	return &DSU[T]{
		parent:  make(map[T]T),
		size:    make(map[T]int),
		maxSize: 0,
	}
}

func (d *DSU[T]) Add(x T) {
	_, ok := d.parent[x]
	if !ok {
		d.parent[x] = x
		d.size[x] = 1
		d.maxSize = max(d.maxSize, 1)
	}
}

func (d *DSU[T]) Find(x T) T {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU[T]) Union(x, y T) bool {
	d.Add(x)
	d.Add(y)

	rootX := d.Find(x)
	rootY := d.Find(y)

	if rootX == rootY {
		return false
	}

	if d.size[rootY] > d.size[rootX] {
		rootX, rootY = rootY, rootX
	}

	d.parent[rootY] = rootX
	d.size[rootX] += d.size[rootY]
	d.maxSize = max(d.maxSize, d.size[rootX])
	return true
}

type DistHeap []Pair

var points []*point

func dist(a, b *point) float64 {

	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (h DistHeap) Len() int {
	return len(h)
}

func (h DistHeap) Less(i, j int) bool {

	return h[i].d > h[j].d
}

func (h DistHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DistHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *DistHeap) Pop() any {
	n := len(*h) - 1
	val := (*h)[n]
	*h = (*h)[:n]
	return val
}

var h *DistHeap

func getClosest(start, end int) {
	if end <= start {
		return
	}
	mid := (start + end) / 2
	dividingX := (points[mid].x + points[mid+1].x) / 2
	getClosest(start, mid)
	getClosest(mid+1, end)
	threshold := math.MaxFloat64
	if h.Len() >= MAX_HEAP_SIZE {
		threshold = (*h)[0].d
	}
	mergePoints1 := make([]*point, 0)
	mergePoints2 := make([]*point, 0)
	for i := mid; i >= start; i-- {
		if float64(dividingX-points[i].x) < threshold {
			mergePoints1 = append(mergePoints1, points[i])
			continue
		}
		break
	}
	for i := mid + 1; i <= end; i++ {
		if float64(-dividingX+points[i].x) < threshold {
			mergePoints2 = append(mergePoints2, points[i])
			continue
		}
		break
	}
	for i := range mergePoints1 {
		for j := range mergePoints2 {
			// 1. Break if Y distance is too big (Optimization from 2D alg)
			if math.Abs(float64(mergePoints1[i].y-mergePoints2[j].y)) >= threshold {
				continue
			}

			// 2. SKIP if Z distance is too big (3D Heuristic)
			// We use Abs because Z is not sorted, so it could be positive or negative difference
			if math.Abs(float64(mergePoints1[i].z-mergePoints2[j].z)) >= threshold {
				continue
			}

			// 3. Only calculate full expensive distance if both Y and Z checks pass
			d := dist(mergePoints1[i], mergePoints2[j])
			if d < threshold {
				heap.Push(h, Pair{mergePoints1[i], mergePoints2[j], d})
				if h.Len() > MAX_HEAP_SIZE {
					heap.Pop(h)
				}
				if h.Len() >= MAX_HEAP_SIZE {
					threshold = (*h)[0].d
				}
			}
		}
	}
}

func main() {
	h = &DistHeap{}
	heap.Init(h)
	inputFile, err := utils.GetInput(8)
	utils.HandleError(err)
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		pointString := scanner.Text()
		pointStringSlice := strings.Split(pointString, ",")
		x, _ := strconv.Atoi(pointStringSlice[0])
		y, _ := strconv.Atoi(pointStringSlice[1])
		z, _ := strconv.Atoi(pointStringSlice[2])
		points = append(points, &point{x, y, z})
	}
	utils.HandleError(scanner.Err())
	slices.SortFunc(points, func(a, b *point) int {
		return cmp.Compare(a.x, b.x)
	})
	getClosest(0, len(points)-1)
	results := make([]Pair, 0, MAX_HEAP_SIZE)
	for h.Len() > 0 {
		results = append(results, heap.Pop(h).(Pair))
	}

	slices.Reverse(results)

	dsu := NewDSU[*point]()

	for i, p := range results {
		dsu.Union(p.p1, p.p2)
		if dsu.maxSize == 1000 {
			fmt.Println(i)
			fmt.Println(p.p1.x * p.p2.x)
			break
		}
	}

	sizes := slices.Collect(maps.Values(dsu.size))
	slices.SortFunc(sizes, func(a, b int) int {
		return -1 * cmp.Compare(a, b)
	})
	product := sizes[0] * sizes[1] * sizes[2]
	fmt.Println(product)
}
