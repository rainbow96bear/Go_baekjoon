package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	averages := make([]int,N)

	for i := 0; i < N; i++ {
		var input string
		fmt.Fscan(reader, &input)
		input = strings.ReplaceAll(input, ".","")
		average, _ := strconv.Atoi(input)
		averages[i] = average
	}

	for peopleCnt:=1 ; peopleCnt<=1000 ; peopleCnt++ {
		if isPossibleCnt(averages, peopleCnt){
			fmt.Fprint(writer, peopleCnt)
			break
		}
	}
}

func isPossibleCnt(averages []int, peopleCnt int) bool {
	for i:=0 ; i<len(averages) ; i++ {
		minTotal := 0
		maxTotal := 10*peopleCnt
		isPossible := false
		for minTotal <= maxTotal{
			midTotal := (minTotal+maxTotal) / 2
			currentAvr := (midTotal*1000)/peopleCnt
			if currentAvr < averages[i]  {
				minTotal = midTotal+1
			}else if currentAvr > averages[i] {
				maxTotal = midTotal-1
			} else {
				if currentAvr > 10*1000 {
					continue
				}
				isPossible = true
				break
			}
		}
		if isPossible == false {
			return isPossible
		}
	}
	return true
}