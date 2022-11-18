package main

import "log"

func SayHello(name string) string {
	return "Hello " + name
}

func main2() {
	//Perimeter(10.0,20.0)
	log.Println(SayHello("Anup"))
}
