package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
type Person struct{
	name string
	state bool
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	members := make(map[string]bool)

	for i:=0 ; i<N ; i++{
		var name, state string
		fmt.Fscan(reader, &name, &state)
		if state == "enter"{
			members[name]=true
		}else if state == "leave"{
			members[name]=false
		}
	}
	rst := []string{}
	for k,v :=range members{
		if v {
			rst = append(rst, k)
		}
	}
	sort.Slice(rst,func(i, j int) bool {
		return rst[i] > rst[j]
	})

	for _, v := range rst {
		fmt.Fprintln(writer,v)
	}
}