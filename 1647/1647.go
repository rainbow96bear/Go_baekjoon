package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
type Load struct {
	from int
	to int
	value int
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()
	N, M := 0, 0
	fmt.Fscanln(reader, &N, &M)

	LoadList := make([]Load, M)
	for i:=0 ; i<M ; i++ {
		from, to, value := 0, 0, 0
		fmt.Fscanln(reader, &from, &to, &value)
		LoadList[i] = Load{from, to, value}
	}
	result := Process(N, LoadList)
	fmt.Fprintln(writer, result)
}

func Process(N int, LoadList []Load) int {

	sort.Slice(LoadList, func(i, j int)bool{return LoadList[i].value < LoadList[j].value})
	Parent := make([]int, N+1)
	for i:=1 ; i<=N ; i++ {
		Parent[i]=i
	}
	total := 0
	LastValue := 0
	for i:=0 ; i<len(LoadList) ; i++ {
		from, to, value := LoadList[i].from, LoadList[i].to, LoadList[i].value
		fromRoot := FindParent(from, Parent)
		toRoot := FindParent(to, Parent)
		if fromRoot != toRoot {
			Parent[fromRoot] = toRoot
			total += value
			LastValue = value
		}
	}
	return total - LastValue
}

func FindParent(num int, Parent []int) int {
	if Parent[num] != num {
		Parent[num] = FindParent(Parent[num], Parent)
		return Parent[num]
	}
	return num
}