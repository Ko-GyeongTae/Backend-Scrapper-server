package main

import (
	"fmt"

	"github.com/Ko-Gyeongtae/learngo/accounts"
)

//#2.0 Bank ToyProject
func main () {
	account := accounts.NewAccount("ko")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}

/*
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

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	//Pointer
	a := 1
	b := &a
	fmt.Println(b)

	//fmt.Println
	fmt.Println("Hello world")

	//function part one
	fmt.Println(multiply(2,2))

	//function part one multi return value
	totalLength, upperName := lenAndUpper("kokt")
	fmt.Println(totalLength, upperName)

	//function part two
	total, up := _lenAndUpper("kokt0203")
	fmt.Println(total, up)

	//for, range
	repeatMe("ko", "lee", "kang", "hyeon")
	result := superAdd(1,2,3,4,5,6,7)
	fmt.Println(result)
	fmt.Println(canIDrink(16))
	names := []string{"ko" ," kt" ,"dal"}

	//Arrays and Slices
	names = append(names, "flynn")
	fmt.Println(names)

	//Maps
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}

	//Structs
	favFood := []string{"kimchi", "ramen"}
	nicolas := person{name:"nico", age:18, favFood:favFood}
	fmt.Println(nicolas)
}
*/