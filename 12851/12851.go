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

	N, K := Input(scanner)
	
	if N == K {
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 1)
		return
	}

	result := Process(N, K)
	for _, v := range result {		
		fmt.Fprintln(writer, v)
	}
}

func Process(N, K int) []int {
	dp := make(map[int]int)
	check := make([]int, 2*100001)
	for i:=0 ; i<=2*100000; i++ {
		dp[i] = 2*1000000
	}
	list := []int{N}
	dp[N] = 0
	check[N] = 1
	for len(list) > 0 {
		now := list[0]
		list = list[1:]

		if now < K {
			if dp[now+1] > dp[now]+1 {
				dp[now+1] = dp[now]+1
				check[now+1] = check[now]
				list = append(list, now+1)
			}else if dp[now+1] == dp[now]+1 {
				check[now+1] += check[now]
			}
			if dp[now*2] > dp[now]+1 {
				dp[now*2] = dp[now]+1
				check[now*2] = check[now]
				list = append(list, now*2)
			}else if dp[now*2] == dp[now]+1 {
				check[now+1] += check[now]
			}
		}
		if now-1 >= 0 {
			if dp[now-1] > dp[now]+1 {
				dp[now-1] = dp[now]+1
				check[now-1] = check[now]
				list = append(list, now-1)
			}else if dp[now-1] == dp[now]+1 {
				check[now-1] += check[now]
			}
		}
	}
	return []int{dp[K], check[K]}
}

func Input(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	input := scanner.Text()

	list := strings.Split(input, " ")
	N, _ := strconv.Atoi(list[0])
	K, _ := strconv.Atoi(list[1])

	return N, K
}