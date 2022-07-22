package main

import (
	"math"
	"sync"
)


const unitDiscoverProcessDivision int = 5000
var primeNumbersCache []int
var wg sync.WaitGroup




func getAllPrimeNumberLowerThan(number int) (primeNumbersSlice [] int){
	discoverPrimes(number)
	for _, v := range primeNumbersCache {
		if v > number {
			continue
		}
		primeNumbersSlice = append(primeNumbersSlice, v)
	}
	return
}


func discoverPrimesUnit(initialNumber int, finalNumber int, maxNumberCached int, wg *sync.WaitGroup) {
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
	defer wg.Done()
}


func discoverPrimes(checkingNumber int) {
	maxNumberCached := max(primeNumbersCache)

	if maxNumberCached < checkingNumber {
		discoverUnits := float64(checkingNumber) / float64(unitDiscoverProcessDivision)

		if int(math.Floor(discoverUnits)) == 0 {
			discoverUnits = 1.0
		}

		for i := int(math.Ceil(discoverUnits)); i >= 1; i--{
			wg.Add(1)
			if i > 1 {
				go discoverPrimesUnit(i * unitDiscoverProcessDivision, (i - 1) * unitDiscoverProcessDivision, maxNumberCached, &wg)
			} else {
				if checkingNumber > unitDiscoverProcessDivision {
					go discoverPrimesUnit(checkingNumber - (unitDiscoverProcessDivision * int(discoverUnits)), 0, maxNumberCached, &wg)
				} else {
					go discoverPrimesUnit(checkingNumber, 0, maxNumberCached, &wg)
				}
			}
		}
		wg.Wait()
	}
}
