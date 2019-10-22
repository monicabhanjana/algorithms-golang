package main

import (
	"custom"
	"fmt"
	"searching"
	"sorting"
	"time"
)

func main() {
	fmt.Println("\nWelcome.")

outerOptionsAgain:
	fmt.Print("\nPlease select an option: \n1) Sorting \n2) Searching \n\nEnter your choice: ")

	var option int
	fmt.Scan(&option)
	fmt.Println()

	switch option {
	case 1:
	innerOptionsAgainSorting:
		fmt.Print("a) Bubble Sort \nb) Selection Sort \n\nEnter your choice: ")

		var option string
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case "a":
			sorting.BubbleSort()
		case "b":
			sorting.SelectionSort()
		default:
			custom.CustomPrint(time.Second, "Invalid Choice! Please try again. \n\n")
			goto innerOptionsAgainSorting
		}
	case 2:
	innerOptionsAgainSearching:
		fmt.Print("a) Linear Search \nb) Binary Search \n\nEnter your choice: ")

		var option string
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case "a":
			searching.LinearSearch()
		case "b":
		default:
			custom.CustomPrint(time.Second, "Invalid Choice! Please try again. \n\n")
			goto innerOptionsAgainSearching
		}
	default:
		custom.CustomPrint(time.Second, "Invalid Choice! Please try again. \n\n")
		goto outerOptionsAgain
	}
}
