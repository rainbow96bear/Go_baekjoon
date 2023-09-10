package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var x1, y1, x2, y2, x3, y3 float64
	fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x3, &y3)
	if (x1 == x2 && x2 == x3) || ((y1-y2)/(x1-x2) == (y2-y3)/(x2-x3)){
		fmt.Fprint(writer, -1)
		return
	}

	side1 := math.Pow((math.Pow(float64(x1-x2), 2.0) + math.Pow(float64(y1-y2), 2.0)) , 0.5)
	side2 := math.Pow((math.Pow(float64(x2-x3), 2.0) + math.Pow(float64(y2-y3), 2.0)) , 0.5)
	side3 := math.Pow((math.Pow(float64(x3-x1), 2.0) + math.Pow(float64(y3-y1), 2.0)) , 0.5)
	
	arr := []float64{side1, side2, side3}
	sort.Float64s(arr)

	fmt.Fprint(writer, float64((arr[2]-arr[0])*2))
}