package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	friendsList := make([][]int, N)
	for i:=0 ; i<N ; i++{
		input, _ := reader.ReadString('\n')
		strings.TrimSpace(input)
		
		for idx, v := range input{
			if v == 'Y'{
				friendsList[i] = append(friendsList[i], idx)
			}
		}
	}
	max := 0
	for i:=0 ; i<N ; i++{
		check := make(map[int]bool, N)
		for j:=0 ; j<len(friendsList[i]) ; j++{
			check[friendsList[i][j]]=true
			friendIdx := friendsList[i][j]
			for k:=0 ; k<len(friendsList[friendIdx]) ; k++{
				check[friendsList[friendIdx][k]]=true
			}
		}
		cnt := 0
		for j:=0 ; j<N ; j++ {
			if check[j] {
				cnt++
			}
		}
		if max < cnt-1{
			max = cnt-1
		}
	}
	fmt.Fprint(writer,max)
}