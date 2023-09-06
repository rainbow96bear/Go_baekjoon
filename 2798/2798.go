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

	cards := make([]int,N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &cards[i])
	}
	sort.Ints(cards)
	max := 0
	for i:=0 ; i<N-2 ; i++ {
		for j:=i+1 ; j<N-1 ; j++{
			for k:=j+1 ; k<N ; k++{
				if cards[i] + cards[j] + cards[k] < M{
					if max < cards[i] + cards[j] + cards[k]{
						max = cards[i] + cards[j] + cards[k]
					}
				}else if cards[i] + cards[j] + cards[k] == M {
					fmt.Fprint(writer, M)
					return
				}
			}
		}
	}
	fmt.Fprint(writer, max)
}