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

	var N, M int
	fmt.Fscan(reader, &N, &M)

	oneLine := make([]int,M)
	sixLine := make([]int,M)
	for i:=0 ; i<M ; i++{
		fmt.Fscan(reader, &sixLine[i], &oneLine[i])
	}

	sort.Ints(oneLine)
	sort.Ints(sixLine)

	if oneLine[0]*6 < sixLine[0] {
		fmt.Fprint(writer, oneLine[0]*N)
	}else {
		if oneLine[0]*(N%6) > sixLine[0]{
			fmt.Fprint(writer, sixLine[0]*((N/6)+1))
		} else {
			fmt.Fprint(writer, sixLine[0]*(N/6)+oneLine[0]*(N%6))
		}
	}
}