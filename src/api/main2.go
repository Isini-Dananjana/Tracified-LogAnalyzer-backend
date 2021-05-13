package main

import (
	"fmt"	
	"io/ioutil"
)


func main() {
	fmt.Println("Hello, 世界")

	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
	  fmt.Println("File reading error", err)
	  return
	}
	fmt.Println("Contents of file:")
	fmt.Println(string(data))
}
