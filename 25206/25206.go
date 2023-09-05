package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	var sum float64 = 0.0
	var totalPoint float64 = 0.0
	totalSub := 20
	for i:=0 ; i<totalSub ; i++{
		var sub string
		var point float64
		var grade string
		fmt.Fscan(reader, &sub, &point, &grade)
		totalPoint+=point
		switch grade {
		case "A+" :
			sum += point*(4.5)
		case "A0" :
			sum += point*(4.0)
		case "B+" :
			sum += point*(3.5)
		case "B0" :
			sum += point*(3.0)
		case "C+" :
			sum += point*(2.5)
		case "C0" :
			sum += point*(2.0)
		case "D+" :
			sum += point*(1.5)
		case "D0" :
			sum += point*(1.0)
		case "F" :
			sum += point*0
		case "P" :
			totalPoint-=point
		}	
	}
	fmt.Fprint(writer, sum/totalPoint)
}