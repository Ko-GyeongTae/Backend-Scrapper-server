package main

import (
	"fmt"
	"strings"
)

func multiply (a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string){
	fmt.Println((words))
}

func _lenAndUpper(name string) (_length int, _uppercase string) {
	defer fmt.Println("I'm done")
	_length = len(name)
	_uppercase = strings.ToUpper(name)
	return 
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper("kokt")
	fmt.Println(totalLength, upperName)
	total, up := _lenAndUpper("kokt0203")
	fmt.Println(total, up)
	repeatMe("ko", "lee", "kang", "hyeon")
}