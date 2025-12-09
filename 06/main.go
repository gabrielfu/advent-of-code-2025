package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(string(content), "\n")
}

type Row []int

type Homework struct {
	rows   []Row
	ops    []string
	length int
}

func ParseLines(lines []string) Homework {
	var rows []Row
	var ops []string

	for _, line := range lines {
		parts := strings.Split(line, " ")
		row := Row{}
		for _, p := range parts {
			if p == "" {
				continue
			}
			if p == "+" || p == "*" {
				ops = append(ops, p)
				continue
			}
			num, _ := strconv.Atoi(p)
			row = append(row, num)
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
	}
	return Homework{
		rows:   rows,
		ops:    ops,
		length: len(ops),
	}
}

func part1(input string) {
	lines := ReadLines(input)
	homework := ParseLines(lines)

	ans := 0
	for i := 0; i < homework.length; i++ {
		op := homework.ops[i]
		if op == "+" {
			rowAns := 0
			for _, row := range homework.rows {
				rowAns += row[i]
			}
			ans += rowAns
		} else if op == "*" {
			rowAns := 1
			for _, row := range homework.rows {
				rowAns *= row[i]
			}
			ans += rowAns
		}
	}
	fmt.Println("Answer is", ans)
}

func SolveProblem(nums []int, ops string) int {
	if ops == "+" {
		sum := 0
		for _, n := range nums {
			sum += n
		}
		return sum
	} else if ops == "*" {
		prod := 1
		for _, n := range nums {
			prod *= n
		}
		return prod
	}
	return 0
}

func part2(input string) {
	lines := ReadLines(input)
	width := len(lines[0])

	ans := 0
	ops := ""
	var nums []int
	for i := 0; i < width; i++ {
		column := ""
		for _, line := range lines {
			if line[i] != ' ' {
				column += string(line[i])
			}
		}

		// reset
		if column == "" {
			ans += SolveProblem(nums, ops)
			ops = ""
			nums = []int{}
			continue
		}

		if column[len(column)-1] == '+' || column[len(column)-1] == '*' {
			ops = string(column[len(column)-1])
			column = column[:len(column)-1]
		}
		num, _ := strconv.Atoi(column)
		nums = append(nums, num)
	}
	ans += SolveProblem(nums, ops)
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
