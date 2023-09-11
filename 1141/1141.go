package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	strGroup := make([]string, N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &strGroup[i])
	}
	relate := make([][]int,N)
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++ {
			if i == j {
				continue
			}
			check := checkInclude(strGroup[i], strGroup[j])
			if check {
				relate[i] = append(relate[i], j)
			}
		}
	}
	answer := 0

	fmt.Fprint(writer, answer)
}

func checkInclude(strA, strB string) bool{
	
	length := len(strA)
	if  len(strB) < length  {
		length = len(strB)
	}

	include := true
	for k:=0 ; k<length ; k++{	
		if strA[k] != strB[k]{
			include = false
			break
		}
	}
	
	return include
}