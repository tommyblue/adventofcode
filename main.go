package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const input = "L3, R2, L5, R1, L1, L2, L2, R1, R5, R1, L1, L2, R2, R4, L4, L3, L3, R5, L1, R3, L5, L2, R4, L5, R4, R2, L2, L1, R1, L3, L3, R2, R1, L4, L1, L1, R4, R5, R1, L2, L1, R188, R4, L3, R54, L4, R4, R74, R2, L4, R185, R1, R3, R5, L2, L3, R1, L1, L3, R3, R2, L3, L4, R1, L3, L5, L2, R2, L1, R2, R1, L4, R5, R4, L5, L5, L4, R5, R4, L5, L3, R4, R1, L5, L4, L3, R5, L5, L2, L4, R4, R4, R2, L1, L3, L2, R5, R4, L5, R1, R2, R5, L2, R4, R5, L2, L3, R3, L4, R3, L2, R1, R4, L5, R1, L5, L3, R4, L2, L2, L5, L5, R5, R2, L5, R1, L3, L2, L2, R3, L3, L4, R2, R3, L1, R2, L5, L3, R4, L4, R4, R3, L3, R1, L3, R5, L5, R1, R5, R3, L1"

	Distance(input)
}

func Distance(input string) {
	currentView := "T"
	rightSequence := []string{"T", "R", "B", "L"}
	leftSequence := []string{"T", "L", "B", "R"}
	totalMove := [2]int{0, 0}

	for _, move := range strings.Split(input, ", ") {
		direction := move[:1]
		length, err := strconv.Atoi(move[1:])
		if err != nil {
			panic("Don't know how long the move is")
		}

		if direction == "R" {
			// rightSequence
			nextIndex := (indexOf(currentView, rightSequence) + 1) % 4
			if nextIndex == -1 {
				panic("Can't find index in rightSequence")
			}
			currentView = rightSequence[nextIndex]
		} else if direction == "L" {
			// leftSequence
			nextIndex := (indexOf(currentView, leftSequence) + 1) % 4
			if nextIndex == -1 {
				panic("Can't find index in leftSequence")
			}
			currentView = leftSequence[nextIndex]
		} else {
			panic("Unknown direction")
		}

		moveToDirection(currentView, length, &totalMove)
	}

	fmt.Printf("Moved of a total of %d\n", abs(totalMove[0])+abs(totalMove[1]))
}

func moveToDirection(direction string, length int, moves *[2]int) {
	// fmt.Printf("Moving %s by %d\n", direction, length)
	if direction == "T" {
		moves[0] += length
	} else if direction == "R" {
		moves[1] += length
	} else if direction == "B" {
		moves[0] -= length
	} else if direction == "L" {
		moves[1] -= length
	} else {
		panic("Unknown direction")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func indexOf(value string, input []string) int {
	for index, v := range input {
		if v == value {
			return index
		}
	}
	return -1
}
