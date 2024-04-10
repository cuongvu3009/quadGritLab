package quadD

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		fmt.Println("x and y must be positive number")
		return
	}

	// Print the top row
	z01.PrintRune('A')
	for i := 1; i < x-1; i++ {
		z01.PrintRune('B')
	}
	if x > 1 {
		z01.PrintRune('C')
	}
	z01.PrintRune('\n')

	// Print the middle part
	for j := 1; j < y-1; j++ {
		z01.PrintRune('B')
		for i := 1; i < x-1; i++ {
			z01.PrintRune(' ')
		}
		if x > 1 {
			z01.PrintRune('B')
		}
		z01.PrintRune('\n')
	}

	// Print the bottom row
	if y > 1 {
		z01.PrintRune('A')
		for i := 1; i < x-1; i++ {
			z01.PrintRune('B')
		}
		if x > 1 {
			z01.PrintRune('C')
		}
		z01.PrintRune('\n')
	}
}

func QuadDTest() {
	// Checking if there are enough arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go quadD <width> <height>")
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
	QuadD(width, height)
}
