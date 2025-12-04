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
	r int
	c int
}

type Grid []string

func (g Grid) At(c Coord) byte {
	return g[c.r][c.c]
}

func (g Grid) AllNeighbors(c Coord) []Coord {
	directions := []Coord{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}
	neighbors := []Coord{}
	for _, d := range directions {
		nr, nc := c.r+d.r, c.c+d.c
		if nr >= 0 && nr < len(g) && nc >= 0 && nc < len(g[0]) {
			neighbors = append(neighbors, Coord{nr, nc})
		}
	}
	return neighbors
}

func (g *Grid) Set(c Coord, value byte) {
	row := []rune((*g)[c.r])
	row[c.c] = rune(value)
	(*g)[c.r] = string(row)
}

func part1(input string) {
	grid := Grid(ReadLines(input))
	ans := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			cursor := Coord{r, c}
			if grid.At(cursor) != '@' {
				continue
			}

			nbRolls := 0
			for _, nb := range grid.AllNeighbors(cursor) {
				nbValue := grid.At(nb)
				if nbValue == '@' {
					nbRolls++
				}
			}
			if nbRolls < 4 {
				ans++
			}
		}
	}
	fmt.Println("Answer is", ans)
}

func part2(input string) {
	grid := Grid(ReadLines(input))
	ans := 0

	for {
		toRemove := make(map[Coord]bool)
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				cursor := Coord{r, c}
				if grid.At(cursor) != '@' {
					continue
				}

				nbRolls := 0
				for _, nb := range grid.AllNeighbors(cursor) {
					nbValue := grid.At(nb)
					if nbValue == '@' {
						nbRolls++
					}
				}
				if nbRolls < 4 {
					toRemove[cursor] = true
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}

		ans += len(toRemove)
		for coord := range toRemove {
			grid.Set(coord, '.')
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
