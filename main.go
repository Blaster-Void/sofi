// Package main provides a simple command-line tool to read and display the contents of a file.
// sofi Alpha
package main

import (
	"fmt"
	"os"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
) // add some color for display

func main() {
	if len(os.Args) < 2 { // check if there are enough arguments
		fmt.Println(colorRed + "Not enough arguments! , Please provide a file name." + colorReset) // add a comment here
		return
	} else if len(os.Args) > 2 { // check if there are too many arguments
		fmt.Println(colorRed + "Too many arguments! , Please provide a file name." + colorReset) // add a comment here
		return
	} else {
		op := os.Args[1]         // get the file name from the command line arguments
		file, err := os.Open(op) // open the file
		if err != nil {          // handle the error if the file cannot be opened
			fmt.Println(colorRed + err.Error() + colorReset) // print the error message
		}
		defer file.Close()         // close the file when the function returns
		data, _ := os.ReadFile(op) // read the file contents
		fmt.Println(string(data))  // print the file contents

	}
}
