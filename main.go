package main

import "io/ioutil"
import "fmt"

func check(e error) {
	if e != nil {
		panic("Error")
	}
}
func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	index := 4 // Index of n. 5
	keyboard := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range data {
		if val != '\n' {
			index = calculateResult(val, index)
		} else {
			fmt.Printf("%d", keyboard[index])
		}
	}
}

func calculateResult(move byte, index int) int {
	if move == 'U' {
		if index-3 < 0 {
			return index
		}
		return index - 3
	} else if move == 'D' {
		if index+3 >= 9 {
			return index
		}
		return index + 3
	} else if move == 'R' {
		if (index+1)%3 == 0 {
			return index
		}
		return index + 1
	} else if move == 'L' {
		if (index)%3 == 0 {
			return index
		}
		return index - 1
	} else {
		panic("Wrong move")
	}
}
