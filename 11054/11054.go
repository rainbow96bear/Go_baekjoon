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
	tempN, _ := strconv.ParseInt(Input(scanner), 10, 64)
	N := int(tempN)
		
	list := strings.Split(Input(scanner), " ")
	numList := []int{}
	for _, v := range list {
		num, _ := strconv.Atoi(v)
		numList = append(numList, num)
	}
	result := Process(N, numList)
	fmt.Fprintln(writer, result)
}

func Process(N int, numList []int) int {
	dpInc := make([]int, N)
	for i := 0; i < N; i++ {
		dpInc[i] = 1
		for j := 0; j < i; j++ {
			if numList[i] > numList[j] && dpInc[i] < dpInc[j]+1 {
				dpInc[i] = dpInc[j] + 1
			}
		}
	}

	dpDec := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		dpDec[i] = 1
		for j := N - 1; j > i; j-- {
			if numList[i] > numList[j] && dpDec[i] < dpDec[j]+1 {
				dpDec[i] = dpDec[j] + 1
			}
		}
	}
	
	answer := 0
	for i:=0 ; i<N ; i++ {
		answer = max(answer, dpInc[i]+dpDec[i]-1)
	}
	return answer
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}