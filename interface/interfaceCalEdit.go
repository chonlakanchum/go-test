package main

import (
	"fmt"
	"time"
)

//base int64
type mathOperator interface {
	Plus() int64
	Minus() int64
	Multi() int64
	Div() int64
}

//base int64
type number struct {
	num1 int64
	num2 int64
}

func (n number) init() number {
	//input Data From Here
	var num1, num2 int64
	fmt.Println("Enter First number : ")
	fmt.Scan(&num1)
	fmt.Println("Enter Second number : ")
	fmt.Scan(&num2)

	return number{num1: num1, num2: num2}
}
func (n number) show() {
	//Show current number
	fmt.Println("=======================")
	fmt.Println(n)
	fmt.Println("=======================")
}
func (n number) Plus() int64 {
	//if Plus can use show Plus Number Result
	fmt.Println("=======================")
	fmt.Println("Plus Result is : ", n.num1+n.num2)
	fmt.Println("=======================")
	return n.num1 + n.num2
}
func (n number) Minus() int64 {
	//if Minus can use show Plus Number Minus
	fmt.Println("=======================")
	fmt.Println("Minus Result is : ", n.num1-n.num2)
	fmt.Println("=======================")
	return n.num1 - n.num2
}
func (n number) Multi() int64 {
	//if Multi can use show Plus Number Multi
	fmt.Println("=======================")
	fmt.Println("Multi Result is : ", n.num1*n.num2)
	fmt.Println("=======================")
	return n.num1 * n.num2
}
func (n number) Div() int64 {
	//if Div can use show Plus Number Div
	if n.num2 != 0 {
		fmt.Println("=======================")
		fmt.Println("Div Result is : ", n.num1/n.num2)
		fmt.Println("=======================")
	} else {
		fmt.Println("num2 connot 0")
		return 0
	}
	return 1
}

func calculator(n mathOperator, c int64) {
	//fmt.Println("++++++++++++++", n)
	switch c {
	case 1:
		n.Plus()
	case 2:
		n.Minus()
	case 3:
		n.Multi()
	case 4:
		n.Div()
	default:
		fmt.Println("Input 1-7!!!")
	}
}

func main() {
	n := number{}
	n = n.init()
	for {
		//show input
		n.show()
		//command
		var command int64
		fmt.Println("Enter operation")
		fmt.Println("1. Plus")
		fmt.Println("2. Minus")
		fmt.Println("3. Multi")
		fmt.Println("4. Div")
		fmt.Println("5. Change Number")
		fmt.Println("6. Swap Number") // num1 num2
		fmt.Println("7. Exit")
		fmt.Scanln(&command)
		if command == 5 {
			n = n.init()
		}
		if command == 6 {
			//n.num1, n.num2 = n.num2, n.num1 //เครื่องมือคณิต สลับ
			n.num1 = n.num1 + n.num2
			n.num2 = n.num1 - n.num2
			n.num1 = n.num1 - n.num2
		}
		if command == 7 {
			break
		}
		calculator(n, command)
		time.Sleep(1 * time.Second)
	}
}
