package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Load struct {
	to int
	time int
}
type WarmHole struct {
	to int
	time int
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	TC := Input(scanner)[0]
	
	for t:=0 ; t<TC ; t++ {
		inputBase := Input(scanner)
		N, M, W := inputBase[0], inputBase[1], inputBase[2]

		LoadList := make([][]Load, N+1)
		for m:=0 ; m<M ; m++ {
			inputLoad := Input(scanner)
			S, E, T := inputLoad[0], inputLoad[1], inputLoad[2]
			LoadList[S] = append(LoadList[S], Load{E, T})
			LoadList[E] = append(LoadList[E], Load{S, T})
		}

		WarmHoleList := make([][]WarmHole, N+1)
		haveWarmHoleList := make(map[int]bool)
		for w:=0 ; w<W ; w++ {
			inputWarmHole := Input(scanner)
			S, E, T := inputWarmHole[0], inputWarmHole[1], inputWarmHole[2]
			WarmHoleList[S] = append(WarmHoleList[S], WarmHole{E, -T})
			haveWarmHoleList[E] = true
		}
		result := Process(N, LoadList, WarmHoleList, haveWarmHoleList)
		fmt.Fprintln(writer, result)
	}
}

func Process(N int, LoadList [][]Load, WarmHoleList [][]WarmHole, haveWarmHoleList map[int]bool) string {
	for i:=1 ; i<=N ; i++ {
		if haveWarmHoleList[i] == true{
			result := FindRoute(i, N, LoadList, WarmHoleList)
			if result < 0 {
				return "YES"
			}
		}
	}
	return "NO"
}

func FindRoute(start, N int, LoadList [][]Load, WarmHoleList [][]WarmHole) int {
	list := []int{start}
	dp := make(map[int]int)
	for i:=1 ; i<=N ; i++ {
		dp[i] = 1<<30
	}
	dp[start] = 0
	for len(list) > 0 {
		now := list[0]
		list = list[1:]

		for _, next := range WarmHoleList[now] {
			if dp[next.to] > dp[now] + next.time {
				dp[next.to] = dp[now] + next.time
				list = append(list, next.to)
			}
		}
		for _, next := range LoadList[now] {
			if dp[next.to] > dp[now] + next.time {
				dp[next.to] = dp[now] + next.time
				list = append(list, next.to)
			}
		}
		if dp[start] < 0 {
			break
		}
	}
	return dp[start]
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