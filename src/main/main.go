package main

import (
	"fmt"
	"sorting"
)

func main() {
	fmt.Print("Welcome. \nPlease select an option: \n1) Sorting \n2) Searching \n\nEnter your choice: ")

	var option int
	fmt.Scan(&option)
	fmt.Println()

	switch option {
	case 1:
		fmt.Print("a) Bubble Sort \nb) Selection Sort \n\nEnter your choice: ")

		var option string
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case "a":
			sorting.BubbleSort()
		case "b":
		default:
			fmt.Print("Invalid Choice! Please try again. \n\n")
		}
	case 2:
		fmt.Print("a) Linear Search \nb) Binary Search \n\nEnter your choice: ")

		var option string
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case "a":
		case "b":
		default:
			fmt.Print("Invalid Choice! Please try again. \n\n")
		}
	default:
		fmt.Print("Invalid Choice! Please try again. \n\n")
	}
}
