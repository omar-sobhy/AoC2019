package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../test.txt")
	check(err)

	totalModuleF, totalAllF := int64(0), int64(0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		check(err)

		moduleF, totalF := fuel(x)
		totalModuleF += int64(moduleF)
		totalAllF += int64(totalF)
	}

	fmt.Println("One:", totalModuleF)
	fmt.Println("Two:", totalAllF)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fuel(mass int) (moduleF int, allF int) {
	moduleF = mass/3 - 2
	allF = 0
	for next := moduleF; next > 0; next = next/3 - 2 {
		allF += next
	}

	return moduleF, allF
}
