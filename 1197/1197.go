package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
type Edge struct {
	from int
	to int
	value int
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	tempInput := Input(scanner)
	V, E := tempInput[0], tempInput[1]
	EdgeList := []Edge{}
	for i:=0 ; i<E ; i++ {
		tempEdge := Input(scanner)
		EdgeList = append(EdgeList, Edge{tempEdge[0], tempEdge[1], tempEdge[2]})
	}
	
	result := Process(V, EdgeList)

	fmt.Fprintln(writer, result)
}

func Process(V int, EdgeList []Edge) int {

	Parent := make([]int, V+1)
	for i:=1 ; i<=V ; i++ {
		Parent[i]=i
	}
	sort.Slice(EdgeList, func(i, j int)bool {return EdgeList[i].value < EdgeList[j].value})
	
	var totalWeight int

	for i:=0 ; i <len(EdgeList) ; i++ {
		fromRoot := FindRoot(EdgeList[i].from, Parent)
		toRoot := FindRoot(EdgeList[i].to, Parent)

		if fromRoot != toRoot {
			Parent[fromRoot] = toRoot
			totalWeight += EdgeList[i].value
		}
	}
	return totalWeight
}

func FindRoot(num int, Parent []int) int {
	if Parent[num] == num {
		return Parent[num]
	}
	Parent[num] = FindRoot(Parent[num], Parent)
	return Parent[num]
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")
	output := []int{}
	
	for _, v := range input {
		num, _ := strconv.Atoi(v)
		output = append(output, num)
	}

	return output
}