package main

import "fmt"

type mathOperator interface {
	Plus() float64
	Minus() float64
	Multi() float64
	Div() float64
}

type number struct {
	num1 float64
	num2 float64
}

func (n number) Plus() float64 {
	return n.num1 + n.num2
}
func (n number) Minus() float64 {
	return n.num1 - n.num2
}
func (n number) Multi() float64 {
	return n.num1 * n.num2
}
func (n number) Div() float64 {
	return n.num1 / n.num2
}

func measure(m mathOperator, n number) {
	i := true
	var getOperation int
	for i {
		fmt.Println("======================================")
		fmt.Println("First Number : ", n.num1, "Second Number : ", n.num2)
		fmt.Println("======================================")
		fmt.Println("Enter operation")
		fmt.Println("1. Plus\n2. Minus\n3. Multi\n4. Div\n5. Exit")
		fmt.Scan(&getOperation)
		switch getOperation {
		case 1:
			fmt.Println("======================================")
			fmt.Println("Plus Result is : ", m.Plus())
			fmt.Println("======================================")
		case 2:
			fmt.Println("======================================")
			fmt.Println("Minus Result is : ", m.Minus())
			fmt.Println("======================================")
		case 3:
			fmt.Println("======================================")
			fmt.Println("Multi Result is : ", m.Multi())
			fmt.Println("======================================")
		case 4:
			fmt.Println("======================================")
			fmt.Println("Div Result is : ", m.Div())
			fmt.Println("======================================")
		case 5:
			fmt.Println("Good Bye!")
			i = false
		}
	}
}

func main() {
	var num1, num2 int

	fmt.Println("Enter First number : ")
	fmt.Scan(&num1)
	fmt.Println("Enter Second number : ")
	fmt.Scan(&num2)
	num := number{float64(num1), float64(num2)}
	//n := n.init(num1, num2)
	measure(num, num)

}
