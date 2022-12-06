package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(input string) int {
	elfs := strings.Split(input, "\n\n")
	caloriesElfCarrying := make([]int, 0)

	for _, elf := range elfs {
		caloriesElfCarrying = append(caloriesElfCarrying, caloriesOnElf(elf))
	}

	return MaxInt(caloriesElfCarrying)
}

func MaxInt(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]

}

func caloriesOnElf(input string) int {
	var calories int

	c := strings.Split(input, "\n")

	for _, i := range c {
		calorie, _ := strconv.ParseInt(i, 10, 64)
		calories += int(calorie)
	}
	return calories
}

func getFileContent(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}
	return string(file)
}

func main() {
	fmt.Println(part1(getFileContent("input.txt")))
}
