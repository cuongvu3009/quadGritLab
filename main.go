// For testing
package main

import (
	"fmt"
	"os"

	"github.com/garif/quad/custom"
	"github.com/garif/quad/quadA"
	"github.com/garif/quad/quadB"
	"github.com/garif/quad/quadC"
	"github.com/garif/quad/quadD"
	"github.com/garif/quad/quadE"
)

func lineBreak() {
	fmt.Println("----------------------------------")
}

func testAll() {

	fmt.Println(`
*****Team Introduction****
   /\_/\
  ( o.o )
   > ^ <
Cuong  Garif  Arivera
**************************
	`)

	lineBreak()

	// test quadA
	fmt.Println("Test QuadA_Program1: ")
	quadA.QuadA(5, 3)
	fmt.Println("")

	fmt.Println("Test QuadA_Program2: ")
	quadA.QuadA(5, 1)
	fmt.Println("")

	fmt.Println("Test QuadA_Program3: ")
	quadA.QuadA(1, 1)
	fmt.Println("")

	fmt.Println("Test QuadA_Program4: ")
	quadA.QuadA(1, 5)
	fmt.Println("")

	fmt.Println("Test QuadA return nothing because x=0 and y=0")
	quadA.QuadA(0, 0)
	fmt.Println("")

	fmt.Println("Test QuadA return nothing because y negative")
	quadA.QuadA(5, -1)
	fmt.Println("")

	fmt.Println("Test QuadA return nothing because x negative")
	quadA.QuadA(-1, 3)
	fmt.Println("")

	fmt.Println("Test QuadA with big x and y")
	quadA.QuadA(10, 8)
	fmt.Println("")

	lineBreak()

	// test quadB
	fmt.Println("Test QuadB_Program1: ")
	quadB.QuadB(5, 3)
	fmt.Println("")

	fmt.Println("Test QuadB_Program2: ")
	quadB.QuadB(5, 1)
	fmt.Println("")

	fmt.Println("Test QuadB_Program3: ")
	quadB.QuadB(1, 1)
	fmt.Println("")

	fmt.Println("Test QuadB_Program4: ")
	quadB.QuadB(1, 5)
	fmt.Println("")

	fmt.Println("Test QuadB return nothing because x=0 and y=0")
	quadB.QuadB(0, 0)
	fmt.Println("")

	fmt.Println("Test QuadB return nothing because y negative")
	quadB.QuadB(5, -1)
	fmt.Println("")

	fmt.Println("Test QuadB return nothing because x negative")
	quadB.QuadB(-1, 3)
	fmt.Println("")

	fmt.Println("Test QuadB with big x and y")
	quadB.QuadB(9, 13)
	fmt.Println("")

	lineBreak()

	// test quadC
	fmt.Println("Test QuadC_Program1: ")
	quadC.QuadC(5, 3)
	fmt.Println("")

	fmt.Println("Test QuadC_Program2: ")
	quadC.QuadC(5, 1)
	fmt.Println("")

	fmt.Println("Test QuadC_Program3: ")
	quadC.QuadC(1, 1)
	fmt.Println("")

	fmt.Println("Test QuadC_Program4: ")
	quadC.QuadC(1, 5)
	fmt.Println("")

	fmt.Println("Test QuadC return nothing because x=0 and y=0")
	quadC.QuadC(0, 0)
	fmt.Println("")

	fmt.Println("Test QuadC return nothing because y negative")
	quadC.QuadC(5, -1)
	fmt.Println("")

	fmt.Println("Test QuadC return nothing because x negative")
	quadC.QuadC(-1, 3)
	fmt.Println("")

	fmt.Println("Test QuadC with big x and y")
	quadC.QuadC(9, 13)
	fmt.Println("")

	lineBreak()

	// test quadD
	fmt.Println("Test QuadD_Program1: ")
	quadD.QuadD(5, 3)
	fmt.Println("")

	fmt.Println("Test QuadD_Program2: ")
	quadD.QuadD(5, 1)
	fmt.Println("")

	fmt.Println("Test QuadD_Program3: ")
	quadD.QuadD(1, 1)
	fmt.Println("")

	fmt.Println("Test QuadD_Program4: ")
	quadD.QuadD(1, 5)
	fmt.Println("")

	fmt.Println("Test QuadD return nothing because x=0 and y=0")
	quadD.QuadD(0, 0)
	fmt.Println("")

	fmt.Println("Test QuadD return nothing because y negative")
	quadD.QuadD(5, -1)
	fmt.Println("")

	fmt.Println("Test QuadD return nothing because x negative")
	quadD.QuadD(-1, 3)
	fmt.Println("")

	fmt.Println("Test QuadD with big x and y")
	quadD.QuadD(9, 13)
	fmt.Println("")

	lineBreak()

	// test quadE
	fmt.Println("Test QuadE_Program1: ")
	quadE.QuadE(5, 3)
	fmt.Println("")

	fmt.Println("Test QuadE_Program2: ")
	quadE.QuadE(5, 1)
	fmt.Println("")

	fmt.Println("Test QuadE_Program3: ")
	quadE.QuadE(1, 1)
	fmt.Println("")

	fmt.Println("Test QuadE_Program4: ")
	quadE.QuadE(1, 5)
	fmt.Println("")

	fmt.Println("Test QuadE return nothing because x=0 and y=0")
	quadE.QuadE(0, 0)
	fmt.Println("")

	fmt.Println("Test QuadE return nothing because y negative")
	quadE.QuadE(5, -1)
	fmt.Println("")

	fmt.Println("Test QuadE return nothing because x negative")
	quadE.QuadE(-1, 6)
	fmt.Println("")

	fmt.Println("Test QuadE with big x and y")
	quadE.QuadE(9, 13)
	fmt.Println("")

	lineBreak()
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "testAll" {
		testAll()
	} else if len(os.Args) > 1 && os.Args[1] == "quadA" {
		quadA.QuadATest()
	} else if len(os.Args) > 1 && os.Args[1] == "quadB" {
		quadB.QuadBTest()
	} else if len(os.Args) > 1 && os.Args[1] == "quadC" {
		quadC.QuadCTest()
	} else if len(os.Args) > 1 && os.Args[1] == "quadD" {
		quadD.QuadDTest()
	} else if len(os.Args) > 1 && os.Args[1] == "quadE" {
		quadE.QuadETest()
	} else {
		custom.Custom()
	}
}
