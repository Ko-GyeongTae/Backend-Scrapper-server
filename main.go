package main

import (
	"fmt"
	"time"
)

func main(){
	go Count("ko")
	go Count("KT")
	Count("Kim")
}

func Count(person string) {
	for i:=0;i<10;i++{
		fmt.Println(person, "is good", i)
		time.Sleep(time.Second)
	}
}