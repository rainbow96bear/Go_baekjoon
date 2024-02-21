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

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)

	for i, _ := range list{
		fmt.Fscanf(reader, "%d", &list[i])
	}
	fmt.Fscanf(reader, "\n")

	calc := make([]int, 4)
	fmt.Fscanf(reader, "%d %d %d %d", &calc[0], &calc[1], &calc[2], &calc[3])

	resultList := make(map[int]bool)

	
	calculation(1, list[0], list, calc, resultList)

	min := 1000000000
	max := -1000000000
	for i, v := range resultList {
		if v {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
	}
	fmt.Fprintln(writer, max)
	fmt.Fprintln(writer, min)
}

func calculation(deep, result int, list, calc []int, resultList map[int]bool){
	if deep == len(list) {
		if resultList[result] == false {
			resultList[result] = true
		}
		return
	}
	for i:=0 ; i<4 ; i++ {
		if calc[i] > 0 {
			calc[i]--
			nextResult := result
			if i == 0 {
				nextResult += list[deep]
			}else if i == 1 {
				nextResult -= list[deep]
			}else if i == 2 {
				nextResult *= list[deep]
			}else if i == 3 {
				if nextResult < 0 {
					nextResult = -((-nextResult) / list[deep])
				}else {
					nextResult /= list[deep]
				}
			}
			calculation(deep+1, nextResult, list, calc, resultList)
			calc[i]++
		}
	}
}