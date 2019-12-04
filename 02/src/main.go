package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../in.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	read := scanner.Text()
	program := opcodes(read)

	result := compute(duplicate(program))
	fmt.Println("One:", result)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			programCopy := duplicate(program)
			programCopy[1] = i
			programCopy[2] = j
			if compute(programCopy) == 19690720 {
				fmt.Println("Two: noun", i, "verb", j, "result", 100*i+j)
			}
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func duplicate(in []int) (ret []int) {
	ret = make([]int, len(in))
	copy(ret, in)
	return
}

func opcodes(in string) (ret []int) {
	tokens := strings.Split(in, ",")
	for _, token := range tokens {
		opcode, err := strconv.Atoi(token)
		check(err)
		ret = append(ret, opcode)
	}

	return
}

func compute(opcodes []int) int {
	for i := 0; i < len(opcodes); i += 4 {
		opcode := opcodes[i]
		if opcode == 99 {
			break
		}

		x := opcodes[i+1]
		y := opcodes[i+2]
		z := opcodes[i+3]
		if opcode == 1 {
			opcodes[z] = opcodes[x] + opcodes[y]
		} else if opcode == 2 {
			opcodes[z] = opcodes[x] * opcodes[y]
		}
	}

	return opcodes[0]
}
