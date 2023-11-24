package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Route struct {
	to int
	length int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	V := Input(scanner)[0]
	Tree := make([][]Route, V+1)
	for i:=0 ; i<V ; i++ {
		list := Input(scanner)
		head := list[0]
		list = list[1:]
		for len(list) > 0 {
			if list[0] == -1 || len(list) < 2{
				break
			}
			Tree[head] = append(Tree[head], Route{list[0], list[1]})
			list = list[2:]
		}
	}
	
	result := Process(Tree, V)

	fmt.Fprintln(writer, result)
}
func Process(Tree [][]Route, V int) int {
	max := 0
	Diameter := make(map[int]int)
	isVisit := make(map[int]bool)
	FindLength(1, V, Tree, Diameter, isVisit)
	for i:=1 ; i<=V ; i++ {
		if max < Diameter[i] {
			max = Diameter[i]
		}
	}
	return max
}

func FindLength(start, V int, Tree [][]Route, Diameter map[int]int, isVisit map[int]bool) int {
	firstMax := 0
	secondMax := 0
	isVisit[start]= true
	for i:=0 ; i<len(Tree[start]) ; i++ {
		to := Tree[start][i].to
		length := Tree[start][i].length
		if isVisit[to] == false {
			result := length + FindLength(to, V, Tree, Diameter, isVisit)
			if result > firstMax {
				firstMax, secondMax = result, firstMax
			}else if result > secondMax {
				secondMax = result
			} 
		}
	}
	Diameter[start] = firstMax + secondMax
	return firstMax
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