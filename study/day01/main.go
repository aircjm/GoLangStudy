package main

import "fmt"

func VarUse() {
	var i = "hello"
	fmt.Println(i)
}

const name = "hello world"

func ConstUse() string {
	//name = "hello world!!"
	return name
}

func main() {
	VarUse()
}
