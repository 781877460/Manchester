package main

import ("fmt"
	"os)

func main(){
	fmt.println("hello",os.Arg[1:])
}

