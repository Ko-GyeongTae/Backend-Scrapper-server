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

func superAdd(numbers ...int) int {
	total := 0
	for i := 0; i< len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
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
	result := superAdd(1,2,3,4,5,6,7)
	fmt.Println(result)
}