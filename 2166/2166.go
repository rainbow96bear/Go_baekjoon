package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N := Input(scanner)[0]
	xList := []int{}
	yList := []int{}
	for i:=0 ; i<N ; i++ {
		input := Input(scanner)
		xList = append(xList, input[0])
		yList = append(yList, input[1])
	}
	area := Calc(xList, yList, N)
	fmt.Fprintf(writer, "%.1f\n", area)
}

func Calc(xList, yList []int, N int) float64 {
	result := int64(0)
	for i:=0 ; i<N ; i++ {
		result += int64(xList[i%N]) * int64(yList[(i+1)%N])
		result -= int64(xList[(i+1)%N]) * int64(yList[i%N])
	}
	
	answer := abs(float64(result)/2)

	return roundToNearestTenth(answer)
}

func roundToNearestTenth(num float64) float64 {
	rounded := int(num*10 + 0.5)
	return float64(rounded) / 10
}

func abs(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	result := []int{}
	for _, v := range input {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}