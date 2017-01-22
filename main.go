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

type triangle struct {
	vertex []int
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	validTriangles := 0
	var tmpTriangles []triangle
	completed := false
	for scanner.Scan() {
		if completed || len(tmpTriangles) == 0 {
			tmpTriangles = make([]triangle, 3)
			completed = false
		}
		var elems []string
		tmpElems := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(tmpElems); i++ {
			if tmpElems[i] != "" {
				elems = append(elems, tmpElems[i])
			}
		}
		for i := 0; i < len(elems); i++ {
			s, err := strconv.Atoi(strings.Trim(elems[i], " "))
			if err == nil {
				tmpTriangles[i].vertex = append(tmpTriangles[i].vertex, s)
			} else {
				panic(err)
			}
			if len(tmpTriangles[i].vertex) == 3 {
				validTriangles += calculateResult(tmpTriangles[i].vertex)
				completed = true
			}
		}
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
