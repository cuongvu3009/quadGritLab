package quadA

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		fmt.Println("x and y must be positive number")
		return
	}

	// Print the top side of the rectangle
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune('o')
		} else {
			z01.PrintRune('-')
		}
	}
	z01.PrintRune('\n')

	// Print the middle part of the rectangle
	for i := 0; i < y-1; i++ {
		for j := 0; j < x; j++ {
			if j == 0 || j == x-1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

	// Print the bottom side of the rectangle
	if y > 1 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadATest() {
	// Checking if there are enough arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go quadA <width> <height>")
		return
	}

	// Parsing the arguments into integers
	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error parsing width:", err)
		return
	}

	height, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error parsing height:", err)
		return
	}

	// Now you have width and height as integers, you can use them as needed.
	// For example, passing them to your QuadA function
	QuadA(width, height)
}
