package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	sex int
	switchNum int
}

var switchList []bool

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscanf(reader, "%d", &input)
		if input == 1 {
			switchList = append(switchList, true)
		}else {
			switchList = append(switchList, false)
		}
	}
	fmt.Fscanf(reader, "\n")

	var num int
	fmt.Fscanf(reader, "%d\n", &num)

	for i:=0 ; i<num ; i++ {
		var sex, switchNum int
		fmt.Fscanf(reader, "%d %d\n", &sex, &switchNum)
		if sex == 1 {
			manProcess(N, switchNum)
		}else {
			womanProcess(N, switchNum)
		}
	}

	for i, v := range switchList {
		if i != 0 && i%20 == 0 {
			fmt.Fprintln(writer)
		}
		if v {
			fmt.Fprintf(writer, "%d ", 1)
		}else {
			fmt.Fprintf(writer, "%d ", 0)
		}
	}
}


func manProcess(N, switchNum int){
	num := switchNum
	for num <= N {
		switchList[num-1] = !switchList[num-1]
		num += switchNum 
	}
}

func womanProcess(N, switchNum int){
	switchNum -= 1
	switchList[switchNum] = !switchList[switchNum]
	for gap := 1 ; switchNum-gap >= 0 && switchNum+gap <N ; gap++ {
		if switchList[switchNum-gap] != switchList[switchNum+gap] {
			return
		}else {
			switchList[switchNum-gap] = !switchList[switchNum-gap]
			switchList[switchNum+gap] = !switchList[switchNum+gap]
		}
	}
}