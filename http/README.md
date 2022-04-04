package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// get the content from Body by inserting this content into a byte slice
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
