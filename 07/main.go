package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

type Coord struct {
	r, c int
}

func (c Coord) Move(dr, dc int) Coord {
	return Coord{r: c.r + dr, c: c.c + dc}
}

type Grid struct {
	rows []string
	h, w int
}

func FromLines(lines []string) Grid {
	return Grid{rows: lines, h: len(lines), w: len(lines[0])}
}

func (g Grid) At(c Coord) byte {
	return g.rows[c.r][c.c]
}

func (g *Grid) Set(c Coord, v byte) {
	row := []byte(g.rows[c.r])
	row[c.c] = v
	g.rows[c.r] = string(row)
}

func (g Grid) InBounds(c Coord) bool {
	return c.r >= 0 && c.r < g.h && c.c >= 0 && c.c < g.w
}

func (g Grid) StartPosition() Coord {
	for r := 0; r < g.h; r++ {
		for c := 0; c < g.w; c++ {
			if g.rows[r][c] == 'S' {
				return Coord{r, c}
			}
		}
	}
	return Coord{-1, -1}
}

func part1(input string) {
	lines := ReadLines(input)
	grid := FromLines(lines)
	start := grid.StartPosition()

	var q []Coord
	q = append(q, start)
	visited := make(map[Coord]bool)
	splits := make(map[Coord]bool)

	for len(q) > 0 {
		// pop
		curr := q[0]
		q = q[1:]

		if !grid.InBounds(curr) || visited[curr] {
			continue
		}
		visited[curr] = true

		// go down
		next := curr.Move(1, 0)
		if !grid.InBounds(next) {
			continue
		}

		// split
		if grid.At(next) == '^' {
			splits[next] = true
			left := next.Move(0, -1)
			right := next.Move(0, 1)

			q = append(q, left)
			q = append(q, right)
		} else {
			q = append(q, next)
		}
	}
	ans := len(splits)

	fmt.Println("Answer is", ans)
}

type Record struct {
	coord Coord
	prev  Coord
}

func part2(input string) {
	lines := ReadLines(input)
	grid := FromLines(lines)
	start := grid.StartPosition()

	var q []Record
	q = append(q, Record{coord: start, prev: Coord{-1, -1}})
	visited := make(map[Record]bool)
	counts := make(map[Coord]int)
	counts[start] = 1

	for len(q) > 0 {
		// pop
		record := q[0]
		q = q[1:]
		curr := record.coord
		prev := record.prev

		if !grid.InBounds(curr) {
			continue
		}
		if visited[record] {
			continue
		}
		visited[record] = true
		counts[curr] += counts[prev]

		// go down
		next := curr.Move(1, 0)
		if !grid.InBounds(next) {
			continue
		}

		// split
		if grid.At(next) == '^' {
			left := next.Move(0, -1)
			right := next.Move(0, 1)

			q = append(q, Record{coord: left, prev: curr})
			q = append(q, Record{coord: right, prev: curr})
		} else {
			q = append(q, Record{coord: next, prev: curr})
		}
	}

	ans := 0
	for coord, count := range counts {
		if coord.r == grid.h-1 {
			ans += count
		}
	}

	fmt.Println("Answer is", ans)
}

func main() {
	isSample := flag.Bool("s", false, "use sample input")
	flag.Parse()

	input := "input.txt"
	if *isSample {
		input = "input_sample.txt"
	}

	var start time.Time
	start = time.Now()
	part1(input)
	fmt.Println("Part 1 finished in:", time.Since(start))

	start = time.Now()
	part2(input)
	fmt.Println("Part 2 finished in:", time.Since(start))
}
