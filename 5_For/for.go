package main

import "fmt"

func main() {
	i:=1
	

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	//continue
	//break

	for i:=range 3 {
		fmt.Println(i)
	}


	if age:=21;age>18{
		fmt.Println("You are an adult")
	}else if age==18{
		fmt.Println("You are 18")
	}else if age>12{
		fmt.Println("You are a teenager")
	}else if age>0{
		fmt.Println("You are a child")
	}
}