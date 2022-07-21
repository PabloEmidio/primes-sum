package main


import (
	"time"
	"strconv"
)

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

func getTimeInterval(initialTime time.Time, finalTime time.Time) string {
	hoursPassed := strconv.Itoa(finalTime.Hour() - initialTime.Hour())
	minutesPassed := strconv.Itoa(finalTime.Minute() - initialTime.Minute())
	secondsPassed := strconv.Itoa(finalTime.Second() - initialTime.Second())

	return hoursPassed + "h" + minutesPassed + "min" + secondsPassed + "s"
}
