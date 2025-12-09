package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ihebu/dsu"
)

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

type Coord struct {
	x, y, z int
}

// squared distance
func (c Coord) Distance(other Coord) int {
	dx := c.x - other.x
	dy := c.y - other.y
	dz := c.z - other.z
	return dx*dx + dy*dy + dz*dz
}

type Node int

type Edge struct {
	from, to Node
	weight   int
}

type Graph struct {
	nodes map[Node]struct{}
	edges []Edge
	v     int
	e     int
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[Node]struct{}),
		edges: []Edge{},
	}
}

func (g *Graph) AddNode(n Node) {
	if _, exists := g.nodes[n]; !exists {
		g.nodes[n] = struct{}{}
		g.v++
	}
}

func (g *Graph) AddEdge(from, to Node, weight int) {
	g.AddNode(from)
	g.AddNode(to)
	g.edges = append(g.edges, Edge{from: from, to: to, weight: weight})
	g.e++
}

func NewCompleteGraph(coords []Coord) *Graph {
	g := NewGraph()
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			dist := coords[i].Distance(coords[j])
			g.AddEdge(Node(i), Node(j), dist)
		}
	}
	return g
}

func ParseLines(lines []string) []Coord {
	var coords []Coord
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		var coord Coord
		fmt.Sscanf(line, "%d,%d,%d", &coord.x, &coord.y, &coord.z)
		coords = append(coords, coord)
	}
	return coords
}

func part1(input string, topK int) {
	lines := ReadLines(input)
	coords := ParseLines(lines)
	graph := NewCompleteGraph(coords)

	// disjoint set
	d := dsu.New()
	for n := range graph.nodes {
		d.Add(n)
	}

	// sort the edges by ascending weight
	edges := graph.edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// connect topK smallest edges
	for i := 0; i < topK; i++ {
		edge := edges[i]
		from := edge.from
		to := edge.to
		if d.Find(from) != d.Find(to) {
			d.Union(from, to)
		}
	}

	sizes := make(map[Node]int)
	for n := range graph.nodes {
		root := d.Find(n).(Node)
		sizes[root]++
	}

	sortedSizes := []int{}
	for _, size := range sizes {
		sortedSizes = append(sortedSizes, size)
	}
	sort.Slice(sortedSizes, func(i, j int) bool {
		return sortedSizes[i] > sortedSizes[j]
	})

	ans := sortedSizes[0] * sortedSizes[1] * sortedSizes[2]
	fmt.Println("Answer is", ans)
}

func part2(input string) {
	lines := ReadLines(input)
	coords := ParseLines(lines)
	graph := NewCompleteGraph(coords)

	// disjoint set
	d := dsu.New()
	for n := range graph.nodes {
		d.Add(n)
	}

	// sort the edges by ascending weight
	edges := graph.edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	ans := 0
	for _, edge := range edges {
		from := edge.from
		to := edge.to
		if d.Find(from) != d.Find(to) {
			d.Union(from, to)
		}

		// see if the whole graph is connected
		root := d.Find(from)
		connected := true
		for n := range graph.nodes {
			if d.Find(n) != root {
				connected = false
				break
			}
		}
		if connected {
			ans = coords[edge.from].x * coords[edge.to].x
			break
		}
	}

	fmt.Println("Answer is", ans)
}

func main() {
	isSample := flag.Bool("s", false, "use sample input")
	flag.Parse()

	input := "input.txt"
	topK := 1000
	if *isSample {
		input = "input_sample.txt"
		topK = 10
	}

	var start time.Time
	start = time.Now()
	part1(input, topK)
	fmt.Println("Part 1 finished in:", time.Since(start))

	start = time.Now()
	part2(input)
	fmt.Println("Part 2 finished in:", time.Since(start))
}
