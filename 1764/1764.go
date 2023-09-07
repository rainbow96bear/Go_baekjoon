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

	Dlist := make([]string, N)
	DBlist := []string{}
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &Dlist[i])
	}
	sort.Strings(Dlist)

	for i:=0 ; i<M ; i++ {
		var input string
		fmt.Fscan(reader, &input)

		index := sort.Search(len(Dlist),func(j int)bool{
			return Dlist[j]>=input
		})
		if index < len(Dlist) && Dlist[index] == input {
			DBlist = append(DBlist, input)
		}
	}
	sort.Strings(DBlist)
	fmt.Fprintln(writer,len(DBlist))
	for _, v := range DBlist {
		fmt.Fprintln(writer, v)
	}
}