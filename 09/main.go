package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/twpayne/go-geos"
)

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

type Coord struct {
	x, y int
}

func Area(c1, c2 Coord) int {
	width := c2.x - c1.x
	height := c2.y - c1.y
	if width < 0 {
		width = -width
	}
	if height < 0 {
		height = -height
	}
	return (width + 1) * (height + 1)
}

func part1(input string) {
	lines := ReadLines(input)
	ans := 0

	var coords []Coord
	for _, line := range lines {
		x := 0
		y := 0
		fmt.Sscanf(line, "%d,%d", &x, &y)
		coords = append(coords, Coord{y: y, x: x})
	}

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			area := Area(coords[i], coords[j])
			if area > ans {
				ans = area
			}
		}
	}

	fmt.Println("Answer is", ans)
}

func NewPolygonFromXY(x1, y1, x2, y2 float64) *geos.Geom {
	coords := [][]float64{
		{x1, y1},
		{x2, y1},
		{x2, y2},
		{x1, y2},
		{x1, y1},
	}
	coordss := [][][]float64{coords}
	return geos.NewPolygon(coordss)
}

func part2(input string) {
	lines := ReadLines(input)
	// fmt.Println(lines)
	ans := 0

	var coords [][]float64
	for _, line := range lines {
		x := 0
		y := 0
		fmt.Sscanf(line, "%d,%d", &x, &y)
		coords = append(coords, []float64{float64(x), float64(y)})
	}
	coords2 := append(coords, coords[0])
	coordss := [][][]float64{coords2}
	polygon := geos.NewPolygon(coordss)

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			p2 := NewPolygonFromXY(coords[i][0], coords[i][1], coords[j][0], coords[j][1])
			if !polygon.Contains(p2) {
				continue
			}
			area := Area(
				Coord{x: int(coords[i][0]), y: int(coords[i][1])},
				Coord{x: int(coords[j][0]), y: int(coords[j][1])},
			)
			if area > ans {
				ans = area
			}
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
