package main


var primeNumbersCache []int


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


func getAllPrimeNumberLowerThan(number int) (primeNumbersSlice [] int){
	discoverPrimes(number)
	for _, v := range primeNumbersCache {
		if v > number {
			break
		}
		primeNumbersSlice = append(primeNumbersSlice, v)
	}
	return
}


func discoverPrimes(checkingNumber int) {
	maxNumberCached := max(primeNumbersCache)

	if maxNumberCached < checkingNumber {
		simpleNotPrimeValidation := func(number int) bool {
			return (number > 2 && number % 2 == 0) || (number > 5 && number % 5 == 0) 
		}

		for checkingNumber > 1 {
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
	}
}
