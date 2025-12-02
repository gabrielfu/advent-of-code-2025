package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

type Id struct {
	Value     int
	Text      string
	Length    int
	LengthOdd bool
}

func FromValue(value int) Id {
	text := strconv.Itoa(value)
	return Id{Value: value, Text: text, Length: len(text), LengthOdd: len(text)%2 == 1}
}

func FromText(text string) Id {
	value := Must(strconv.Atoi(text))
	return Id{Value: value, Text: text, Length: len(text), LengthOdd: len(text)%2 == 1}
}

type Range struct {
	Start Id
	End   Id
}

func ReadInput(filename string) []Range {
	content := Must(os.ReadFile(filename))
	var ranges []Range
	inputs := strings.Split(strings.TrimSpace(string(content)), ",")
	for _, part := range inputs {
		bounds := strings.Split(part, "-")
		ranges = append(ranges, Range{
			Start: FromText(bounds[0]),
			End:   FromText(bounds[1]),
		})
	}
	return ranges
}

// Part 1
func SumOfInvalidIds(r Range) int {
	// if start & end both have the same odd length, no invalid IDs
	if r.Start.Length == r.End.Length && r.Start.LengthOdd {
		return 0
	}

	sum := 0
	cursor := r.Start
	for {
		if cursor.Value > r.End.Value {
			break
		}
		if cursor.LengthOdd {
			cursor = FromValue(int(math.Pow(10, float64(cursor.Length))))
			continue
		}
		firstHalf := cursor.Text[:cursor.Length/2]
		secondHalf := cursor.Text[cursor.Length/2:]
		if firstHalf == secondHalf {
			sum += cursor.Value
			nextFirstHalf := strconv.Itoa(Must(strconv.Atoi(firstHalf)) + 1)
			cursor = FromText(nextFirstHalf + nextFirstHalf)
			continue
		} else if secondHalf < firstHalf {
			cursor = FromText(firstHalf + firstHalf)
			continue
		} else {
			nextFirstHalf := strconv.Itoa(Must(strconv.Atoi(firstHalf)) + 1)
			cursor = FromText(nextFirstHalf + nextFirstHalf)
			continue
		}
	}
	return sum
}

func part1(input string) {
	ranges := ReadInput(input)
	ans := 0

	for _, r := range ranges {
		ans += SumOfInvalidIds(r)
	}

	fmt.Println("Answer is", ans)
}

// Part 2
func CheckInvalidId(v int) bool {
	str := strconv.Itoa(v)
	length := len(str)
	for i := 1; i <= length/2; i++ {
		if length%i == 0 && str[i:] == str[:length-i] {
			return true
		}
	}
	return false
}

func part2(input string) {
	ranges := ReadInput(input)
	ans := 0

	for _, r := range ranges {
		for v := r.Start.Value; v <= r.End.Value; v++ {
			if CheckInvalidId(v) {
				ans += v
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
