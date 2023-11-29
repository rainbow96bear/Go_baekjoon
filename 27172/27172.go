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

	N := 0
	fmt.Fscanf(reader, "%d\n", &N)
	list := make([]int, N)
	check := make(map[int]bool)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &list[i])
		check[list[i]] = true
	}
	result := Process(N, list, check)
	for _, v := range list {
		fmt.Fprintf(writer, "%d ", result[v])
	}
}

func Process(N int, list []int, check map[int]bool) map[int]int {
	dp := make(map[int]int)
	for i:=0 ; i<N ; i++ {
		for j:=list[i] * 2 ; j<=1000000 ; j += list[i] {
			if check[j] == true {
				dp[list[i]] ++ 
				dp[j]--
			}	
		}
	}
	return dp
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	result := []int{}

	for _,v := range input {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}

	return result
}