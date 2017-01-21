package main

import "io/ioutil"
import "fmt"
import "math"

func check(e error) {
	if e != nil {
		panic("Error")
	}
}
func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	index := 10 // Index of n. 5
	keyboard := []int{
		0, 0, 1, 0, 0,
		0, 2, 3, 4, 0,
		5, 6, 7, 8, 9,
		0, 10, 11, 12, 0,
		0, 0, 13, 0, 0}
	for _, val := range data {
		if val != '\n' {
			index = calculateResult(val, index, keyboard)
		} else {
			fmt.Printf("%d ", keyboard[index])
		}
	}
}

func calculateResult(move byte, index int, keyboard []int) int {
	length := len(keyboard)
	increment := int(math.Sqrt(float64(length)))
	if move == 'U' {
		if index-increment < 0 || keyboard[index-increment] == 0 {
			return index
		}
		return index - increment
	} else if move == 'D' {
		if index+increment >= length || keyboard[index+increment] == 0 {
			return index
		}
		return index + increment
	} else if move == 'R' {
		if (index+1)%increment == 0 || keyboard[index+1] == 0 {
			return index
		}
		return index + 1
	} else if move == 'L' {
		if (index)%increment == 0 || keyboard[index-1] == 0 {
			return index
		}
		return index - 1
	} else {
		panic("Wrong move")
	}
}
