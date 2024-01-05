package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	CraneList := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &CraneList[i])
	}
	fmt.Fscanf(reader, "\n")

	var M int
	fmt.Fscanf(reader, "%d\n", &M)

	BoxList := make([]int, M)
	for i:=0 ; i<M ; i++ {
		fmt.Fscan(reader, &BoxList[i])
	}
	fmt.Fscanf(reader, "\n")

	result := FindMinTime(N, M, CraneList, BoxList)
	fmt.Fprintln(writer, result)
}

func FindMinTime(N, M int, CraneList, BoxList []int) int {
	sort.Slice(CraneList, func(i, j int) bool {return CraneList[i]>CraneList[j]})
	sort.Slice(BoxList, func(i, j int) bool {return BoxList[i]>BoxList[j]})
	if CraneList[0] < BoxList[0] {
		return -1
	}
	time := 0

	for len(BoxList) > 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < len(BoxList); j++ {
				if CraneList[i] >= BoxList[j] {
					BoxList = append(BoxList[:j], BoxList[j+1:]...)
					break
				}
			}
		}
		time++
	}

	return time
}