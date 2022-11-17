package main

import "fmt"

func main() {
	var param1, param2, area float64
	var shape string
	fmt.Println("Please enter shape name")
	fmt.Scan(&shape)
	fmt.Println("Please enter value of first side")
	fmt.Scan(&param1)
	fmt.Println("Please enter value of second side")
	fmt.Scan(&param2)
	calculateArea(shape,param1,param2)
}

/*
math operation

	func getResult(num1 int, num2 int, operation string) int {
		var result int
		switch operation {
		case "add":
			result = num1 + num2
			break
		case "aubtract":
			result = num1 - num2
			break
		case "multiplied by":
			result = num1 * num2
			break
		case "divided by":
			result = num1 / num2
			break
		}
		return result
	}

	func calculateFactorial(num int) int {
		var numFact int = 1
		for num > 0 {
			numFact = numFact * num
			num--
		}
		return numFact
	}

	func isPrime(num int) bool {
		if num == 0 || num == 1 {
			return false
		}
		if num == 2 {
			return true
		}
		for i := 2; i < num/2; i++ {
			if num/i == 0 {
				return false
			}
		}
		return true
	}

	func isArmstrong(num int) bool {
		var sum float64
		digits := float64(len(strconv.Itoa(num)))
		for _, ch := range strconv.Itoa(num) {
			digit, _ := strconv.Atoi(string(ch))

			sum += math.Pow(float64(digit), digits)
		}
		if int(sum) == num {
			return true
		}
		return false
	}

func printFibonacciSeries(num1 int, num2 int, num3 int, length int) {
	for i := 0; i < length; i++ {
		num1 = num2
		num2 = num3
		num3 = num1 + num2
		fmt.Print(num1, " ")
	}
}*/

/*
String manipulation

	func stringReverse(s string) string {
		//with string methods
		var revString string
		for i := len(s) - 1; i >= 0; i-- {
			revString += string(s[i])
		}
		return revString
	}

	func isPalindrome(s string) bool {
		result := stringReverse(s)
		if s == result {
			return true
		} else {
			return false
		}
	}

func findDuplicates(s string) {
	var myMap = make(map[string]int)
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		value, ok := myMap[string(chars[i])]
		if ok {
			myMap[string(chars[i])] = value + 1
		} else {
			myMap[string(chars[i])] = 1
		}
	}
	fmt.Println(myMap)

}*/

func calculateArea(shape string, param1 float64, param2 float64) {
switch shape{
	case "Circle":
		area = 3.1415 * param1*param1
		
	case "Rectangle":
		area = param1*param2
		
	case "square":
		area = param1*param1
		
	case "Triangle":
		area = 0.5 * param1 * param2		
}
fmt.println(area)
}
