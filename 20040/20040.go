package main

import (
	"bufio"
	"fmt"
	"os"
)
type line struct {
	s int
	e int
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	list := []line{}
	for i:=0 ; i<M ; i++ {
		var s, e int
		fmt.Fscanf(reader, "%d %d\n", &s, &e)
		list = append(list, line{s, e})
	}
	result := Process(N, M, list)
	fmt.Fprintln(writer, result)
}

func Process(N, M int, list []line) int {
	parent := make([]int,N)
	for i:=0 ; i<N ; i++ {
		parent[i]=i
	}

	for i:=0 ; i<M ; i++ {
		s, e := list[i].s, list[i].e
		Sparent := FindParent(s, parent)
		Eparent := FindParent(e, parent)
		if Sparent == Eparent{
			return i+1
		}
		if Sparent < Eparent{
			parent[Sparent] = Eparent
		}else {
			parent[Eparent] = Sparent
		}
	}
	return 0
}

func FindParent(num int, parent []int) int {
	if parent[num] != num {
		parent[num] = FindParent(parent[num], parent)
		return parent[num]
	}
	return num
}
