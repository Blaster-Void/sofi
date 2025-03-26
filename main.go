// Package provides a simple command-line tool to read and display the contents of a file.
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
	} else if len(os.Args) > 2 && os.Args[1] != "-m" { // check if there are too many arguments
		for _, arg := range os.Args { // iterate over the arguments
			file, err := os.Open(arg) // open the file
			if err != nil {           // handle the error if the file cannot be opened
				fmt.Println(colorRed + err.Error() + colorReset) // print the error message
			}
			defer file.Close()          // close the file when the function returns
			data, _ := os.ReadFile(arg) // read the file contents
			fmt.Println(string(data))   // print the file contents
		}
	} else if os.Args[1] == "." { // list all files in the current directory
		wd, _ := os.Getwd()          // get the current working directory
		files, _ := os.ReadDir(wd)   // read the directory contents
		for _, file := range files { // iterate over the files in the directory
			file, err := os.Open(file.Name()) // open the file
			if err != nil {                   // handle the error if the file cannot be opened
				fmt.Println(colorRed + err.Error() + colorReset) // print the error message
			}
			defer file.Close()                                                                                                                            // close the file when the function returns
			data, _ := os.ReadFile(file.Name())                                                                                                           // read the file contents
			fmt.Println(colorYellow, file.Name(), colorReset, colorGreen, "\nstart->(\n", colorReset, string(data), colorGreen, "\n)\n<-end", colorReset) // print the file contents
		}
	} else if os.Args[1] == "-m" { // multi args reading file
		for _, arg := range os.Args[2:] { // iterate over the arguments
			file, err := os.ReadFile(arg)
			if err != nil {
				fmt.Println(colorRed + err.Error() + colorReset)
			}
			fmt.Println(colorYellow, arg, colorReset, colorGreen, "\nstart->(\n", colorReset, string(file), colorGreen, "\n)<-end", colorReset) // print the file contents
		}

	} else if len(os.Args) == 2 && os.Args[1] != "-m" {
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
