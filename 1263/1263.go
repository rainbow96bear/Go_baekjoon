package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type schedule struct {
	workTime int
	limitTime int
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)
	
	schedules := make([]schedule, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d %d\n", &schedules[i].workTime, &schedules[i].limitTime)
	}

	sort.Slice(schedules, func(i, j int)bool{return schedules[i].limitTime > schedules[j].limitTime})
	maxTime := schedules[0].limitTime

	for i:=0 ; i<N ; i++ {
		if schedules[i].limitTime < maxTime {
			maxTime = schedules[i].limitTime-schedules[i].workTime
		}else {
			maxTime -= schedules[i].workTime
		}
		if maxTime < 0 {
			maxTime = -1
			break
		}
	}
	fmt.Fprintln(writer, maxTime)
}