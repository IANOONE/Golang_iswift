package main

import (
	"exercise/helpers"
	"os"
)

func main() {
		
	// get argument
	var argsRaw = os.Args

	// check argument
	argTerminal := helpers.CheckArg(argsRaw...)

	//fetch data
	 data := helpers.FetchData(argTerminal)
	
	// print result
	helpers.PrintResult(data)
}

