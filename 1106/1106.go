package main

import (
	"bufio"
	"fmt"
	"os"
)

type Ads struct {
	cost int
	people int
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var C, N int
	fmt.Fscanf(reader, "%d %d\n", &C, &N)

	AdsList := []Ads{}
	for i:=0 ; i<N ; i++ {
		var cost, people int
		fmt.Fscanf(reader, "%d %d\n", &cost, &people)
		AdsList = append(AdsList, Ads{cost, people})
	}
	result := FindMinCost(C,N,AdsList)
	fmt.Fprintln(writer, result)
}

func FindMinCost(C, N int, AdsList []Ads) int {
	dp := make([]int, C+101)
	for i:=1 ; i<=C+100 ; i++ {
		dp[i] = 1<<30
	}

	for _, ads := range AdsList {
		cost, people := ads.cost, ads.people
		for p:=people ; p<=C+100 ; p++ {
			dp[p] = min(dp[p], dp[p-people]+cost)
		}
	}
	minCost := 1<<30
	for i:=C ; i<=C+100 ; i++ {
		minCost = min(minCost, dp[i])
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}