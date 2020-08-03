package main 

import (
	"runtime"
	"fmt"
)

func main(){
	go sayHello()
	go sayWorld()
	var str string
	fmt.Scan(&str)
}

func sayHello(){
	for i := 0; i < 10; i++{
		fmt.Print("Hello")
		runtime.Gosched()
	}
}

func sayWorld(){
	for i:= 0; i < 10; i++{
		fmt.Println("World")
		runtime.Gosched()
	}
}