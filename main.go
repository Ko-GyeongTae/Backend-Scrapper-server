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

func canIDrink(age int) bool {
	switch age {
	case 16:
		fmt.Println("16살 입니다")
	case 17:
		fmt.Println("17살입니다")
	}
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	a := 1
	b := &a
	fmt.Println(b)
	fmt.Println("Hello world")
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper("kokt")
	fmt.Println(totalLength, upperName)
	total, up := _lenAndUpper("kokt0203")
	fmt.Println(total, up)
	repeatMe("ko", "lee", "kang", "hyeon")
	result := superAdd(1,2,3,4,5,6,7)
	fmt.Println(result)
	fmt.Println(canIDrink(16))
	names := []string{"ko" ," kt" ,"dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}