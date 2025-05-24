package main

import "fmt"


func main() {	

	messaege := make(chan string)


	messaege <- "Hello World"


	msg:=<- messaege

	fmt.Println(msg)
}