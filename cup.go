package main

import (
	"fmt"
)

type cupSet struct {
	cup1 bool
	cup2 bool
	cup3 bool
}

func main() {
	cup := cupSet{}
	var coin string
	i := true
	for i {
		fmt.Print("Enter cup :")
		fmt.Scan(&coin)
		switch coin {
		case "A":
			cup.cup1 = true
			i = false
		case "B":
			cup.cup2 = true
			i = false
		case "C":
			cup.cup3 = true
			i = false
		default:
			fmt.Println("Input A B C!!!")
		}
	}
	//fmt.Println(cup)

	var coins string
	fmt.Print("Swap cup :")
	fmt.Scan(&coins)
	for i := 0; i < len(coins); i++ {
		//fmt.Println(string(coins[i]))
		switch string(coins[i]) {
		case "A":
			temp := cup.cup1
			cup.cup1 = cup.cup2
			cup.cup2 = temp
			//fmt.Println(cup)
		case "B":
			temp := cup.cup2
			cup.cup2 = cup.cup3
			cup.cup3 = temp
			//fmt.Println(cup)
		case "C":
			temp := cup.cup3
			cup.cup3 = cup.cup1
			cup.cup1 = temp
			//fmt.Println(cup)
		default:
			fmt.Println("swap A B C!!!")
		}
	}
	//fmt.Println("real ", cup)
	switch {
	case cup.cup1 == true:
		fmt.Println("coin in cup : A")
	case cup.cup2 == true:
		fmt.Println("coin in cup : B")
	case cup.cup3 == true:
		fmt.Println("coin in cup : C")
	}
}
