package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	paths := readLines("../in.txt")
	m := make(map[Point]int)
	for i, path := range paths {
		mask := 1 << i
		mapPath(m, path, mask)
	}

	min := math.MaxInt32
	for k, v := range m {
		if k.x == 0 && k.y == 0 {
			continue
		}
		if v == (1 | 1<<1) {
			d := distance(k)
			if d < min {
				min = d
			}
		}
	}

	fmt.Println("One:", min)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(path string) (lines []string) {
	f, err := os.Open(path)
	defer f.Close()
	check(err)

	lines = make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func mapPath(m map[Point]int, path string, mask int) {
	moves := strings.Split(path, ",")
	pos := Point{0, 0}
	for _, m_ := range moves {
		direction := m_[0]
		magnitude, err := strconv.Atoi(m_[1:])
		check(err)
		move(m, &pos, direction, magnitude, mask)
	}

}

func move(m map[Point]int, pos *Point, direction byte, magnitude int, mask int) {
	for i := 0; i < magnitude; i++ {
		m[*pos] |= mask
		if direction == 'U' {
			pos.y++
		} else if direction == 'D' {
			pos.y--
		} else if direction == 'R' {
			pos.x++
		} else {
			pos.x--
		}
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func distance(point Point) int {
	return abs(point.x) + abs(point.y)
}
