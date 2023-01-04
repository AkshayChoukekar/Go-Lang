package main

func SayHello(name string) string {
	return "Hello " + name
}

func CalculateSquare(x int) int {
	return x * x
}

func main() {
	//Perimeter(10.0,20.0)
	println(SayHello("Anup"))
	println(CalculateSquare(2))
}
