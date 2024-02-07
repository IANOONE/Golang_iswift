package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// input
	fmt.Print("Enter number : ")
	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')
	// clean string
	input = strings.TrimSuffix(input, "\r\n")

	result, err := strconv.Atoi(input)
	
	if err != nil {
		fmt.Println("This is not a number, Please input a number!")
		main()
	}


	for i := 1; i <= result; i++ {
		// square
		square := math.Floor(math.Sqrt(float64(i)))
		compareSquare := math.Pow(float64(square), 2) == float64(i)

		// cube
		cube := math.Floor(math.Cbrt(float64(i)))
		compareCube := math.Pow(float64(cube), 3) == float64(i)

		switch {
		case compareSquare && compareCube :
			fmt.Println("SquareCube")
		case compareSquare :
			fmt.Println("Square")
		case compareCube :
			fmt.Println("Cube") 
		default :
		fmt.Println(i)
		}		 
	}



}
