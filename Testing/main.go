package main
import "log"

func Perimeter(f1, f2 float64) float64 {
	return 2 * (f1 + f2)
}

func SayHello(name string) string {
	return "Hello "+name
}

func main() {
	//Perimeter(10.0,20.0)
	log.Println(SayHello("Anup"))
}
