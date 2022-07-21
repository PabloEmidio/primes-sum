package main


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
