package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic("Error")
	}
}
func main() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	validTriangles := 0

	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), " ")
		var vertex []int
		for i := 0; i < len(elems); i++ {
			if s, err := strconv.Atoi(strings.Trim(elems[i], " ")); err == nil {
				vertex = append(vertex, s)
			}
		}
		validTriangles += calculateResult(vertex)
	}
	fmt.Printf("Found %d valid triangles\n", validTriangles)
}

func calculateResult(vertex []int) int {
	res := 0
	if vertex[0]+vertex[1] > vertex[2] && vertex[1]+vertex[2] > vertex[0] && vertex[0]+vertex[2] > vertex[1] {
		res = 1
	}
	return res
}
