package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Waiting for messages")
	for i := 0; i< len(people); i++ {
		fmt.Println("Received this message:", <-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}

/*
func Count(person string) {
	for i:=0;i<10;i++{
		fmt.Println(person, "is good", i)
		time.Sleep(time.Second)
	}
}
*/