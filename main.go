package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)


func printResult(number int){
	initialTime := time.Now()
	primesSlice := getAllPrimeNumberLowerThan(number)
	finalTime := time.Now()

	alreadyPrinted := func (valueOne int, valueTwo int) bool {
		return valueTwo > valueOne
	}

	for _, v1 := range primesSlice{
		for _, v2 := range primesSlice {
			if v1 + v2 == number  && !alreadyPrinted(v1, v2){
				fmt.Println(strconv.Itoa(v1) + " + " + strconv.Itoa(v2) + " = " + strconv.Itoa(number))
			}
		}
	}

	timeInterval := getTimeInterval(initialTime, finalTime)
	lenPrimes := len(primesSlice)

	fmt.Printf("\nIn %s, it's was found %d primes Number between 2 and %d\n", timeInterval, lenPrimes, number)

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
