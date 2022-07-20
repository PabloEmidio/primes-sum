package main

import (
	"fmt"
	"log"
	"strconv"
)


func printResult(number int){
	primeNumbers := getAllPrimeNumberLowerThan(number)

	alreadyPrinted := func (valueOne int, valueTwo int) bool {
		return valueTwo > valueOne
	}

	for _, v1 := range primeNumbers{
		for _, v2 := range primeNumbers {
			if v1 + v2 == number  && !alreadyPrinted(v1, v2){
				fmt.Println(strconv.Itoa(v1) + " + " + strconv.Itoa(v2) + " = " + strconv.Itoa(number))
			}
		}
	}
}


func main(){
	var number int

	for {
		fmt.Print("Type a number: ")
		_, err := fmt.Scanf("%d", &number)

		if err != nil {
			log.Fatal(err)
		} else if number == 0 {
			break
		} else{
			printResult(number)
		}
	}
	fmt.Println("Program ended!")
}
