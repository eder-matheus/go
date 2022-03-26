// package == project == workspace
// it informs that the code in this file is part of the project 'main'
// the package named 'main' is the one that will creates a executable file
// other names creates 'reusable' packages --> reusable logic, libraries, etc
package main

// import gives access to code written in other package
// similar to C/C++ include
// fmt is a kind of std library for go
import "fmt"

// for a executable package, it is necessary to have a main function
// func keyword is used to declare a function
// remaining structure is similar to C/C++ (function name followed by
// parenthesis and code inside brackets)
func main() {
	fmt.Println("Hi there!")
}

// this code skeleton is the base for any other implementation
// define the package, import the libraries and declare/implement functions
