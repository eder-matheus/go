package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	io.Copy(os.Stdout, file)
}
