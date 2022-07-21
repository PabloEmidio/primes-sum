package main

import "math"


const unitDiscoverProcessDivision int = 10000
var primeNumbersCache []int
var discoverProcessDone bool = true


func max(iter []int) (greaterValue int ){
	for i, value := range iter {
		if i == 0 {
			greaterValue = value
		} else{
			if value > greaterValue {
				greaterValue = value
			}
		}
	}
	return
}



func discoverPrimesUnit(initialNumber int, finalNumber int, maxNumberCached int) {
	simpleNotPrimeValidation := func(number int) bool {
		return (number > 2 && number % 2 == 0) || (number > 5 && number % 5 == 0)
	}

	checkingNumber := initialNumber
	for checkingNumber > finalNumber && checkingNumber > maxNumberCached {
		var isPrime bool = true
		if simpleNotPrimeValidation(checkingNumber){
			isPrime = false
		} else {
			for i := 2; i < checkingNumber; i++ {
				if checkingNumber % i == 0 {
					isPrime = false
				}
			}
		}
		if isPrime {
			primeNumbersCache = append(primeNumbersCache, checkingNumber)
		}
		checkingNumber--
	}

	if finalNumber < maxNumberCached || finalNumber == 0 {
		discoverProcessDone = true
	}
}


func discoverPrimes(checkingNumber int) {
	maxNumberCached := max(primeNumbersCache)

	if maxNumberCached < checkingNumber {
		discoverUnits := float64(checkingNumber) / float64(unitDiscoverProcessDivision)

		if int(discoverUnits) == 0 {
			discoverPrimesUnit(checkingNumber, 0, maxNumberCached)
		} else{
			discoverProcessDone = false

			for i := int(math.Ceil(discoverUnits)); i >= 1; i--{
				if i > 1 {
					go discoverPrimesUnit(i * unitDiscoverProcessDivision, (i - 1) * unitDiscoverProcessDivision, maxNumberCached)
				} else {
					go discoverPrimesUnit(checkingNumber - (unitDiscoverProcessDivision * int(discoverUnits)), 0, maxNumberCached)
				}
			}

			for {
				if discoverProcessDone { break }
			}
		}
	}
}
