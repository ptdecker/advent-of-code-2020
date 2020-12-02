/*
 * Advent of Code - 2020
 */

package main

import (
	"fmt"
	"strings"
)

func tryProblem(name string, actual int, expected int) {
	if expected == actual {
		fmt.Printf("Problem %s answer: %v (pass)\n", name, actual)
	} else {
		fmt.Printf("%s answer: %v (fail, expected: %v)\n", name, actual, expected)
	}
}

func runProblems() {
	tryProblem("01-A", problem01A("./data/day01.txt"), 3297866)
	tryProblem("01-B", problem01B("./data/day01.txt"), 4943923)
}

func main() {

mainloop:
	for {
		val := prompt("Select an option--(R)un Problems, (Q)uit:", ">")
		switch strings.ToUpper(val) {
		case "R":
			runProblems()
		case "Q":
			break mainloop
		default:
			fmt.Println("Please enter either an 'R', or 'Q'")
		}
	}
}
