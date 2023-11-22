package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	prev int
	cost int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N := Input(scanner)[0]
	M := Input(scanner)[0]

	BusList := make(map[int]map[int]int)
	CityList := []int{}
	check := make(map[int]bool)
	for i:=0 ; i<M ; i++ {
		input := Input(scanner)
		from := input[0]
		to := input[1]
		cost := input[2]
		if check[from] == false {
			BusList[from] = make(map[int]int)
			check[from] = true
			CityList = append(CityList, from)
		}
		if check[to] == false {
			BusList[to] = make(map[int]int)
			check[to] = true
			CityList = append(CityList, to)
		}
		if BusList[from][to] == 0 || BusList[from][to] > cost {
			BusList[from][to] = cost
		}
	}
	route := Input(scanner)
	start := route[0]
	end := route[1]
	result := Process(N, start, end, BusList, CityList)

	fmt.Fprintln(writer, result[end].cost)
	Route := []int{end}
	Prev := end
	for Prev != start {
		Prev = result[Prev].prev
		Route = append(Route, Prev)
	}
	fmt.Fprintln(writer, len(Route))
	for i:= len(Route)-1 ; i>=0 ; i-- {
		fmt.Fprintf(writer, "%d ", Route[i])
	}
}

func Process(N, start, end int, BusList map[int]map[int]int, CityList []int) map[int]Node {
	list := []int{start}
	dp := make(map[int]Node)

	for i:=0 ; i<len(CityList) ; i++ {
		dp[CityList[i]] = Node{start, 1<<30}
	}
	dp[start] = Node{start, 0}

	for len(list) > 0 {
		now := list[0]
		list = list[1:]
		for i, v := range BusList[now] {
			to := i
			cost := v
			if dp[to].cost > dp[now].cost+cost {
				dp[to] = Node{now, dp[now].cost+cost}
				list = append(list, to)
			}
		}
	}
	return dp
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	result := make([]int, len(input))

	for i, v := range input {
		result[i], _  = strconv.Atoi(v)
	}

	return result
}