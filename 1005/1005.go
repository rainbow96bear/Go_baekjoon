package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	rx := bufio.NewReader(os.Stdin)
	wx := bufio.NewWriter(os.Stdout)
	
	defer wx.Flush()
	
	var testNum int
	fmt.Fscan(rx, &testNum)
	for i:=0 ; i<testNum ; i++ {
		var buildNum, ruleNum int
		var winBuild int
		fmt.Fscan(rx, &buildNum, &ruleNum)
		buildTime := make([]int,buildNum+1)
		buildRule := make([][]int,buildNum+1)
		for j := 1 ; j<=buildNum ; j++{
			fmt.Fscan(rx,&buildTime[j])
		}
		for k:=1 ; k<=ruleNum ; k++{
			var key, value int
			fmt.Fscan(rx, &value, &key)
			buildRule[key]=append(buildRule[key],value)
		}
		fmt.Fscan(rx,&winBuild)

		fmt.Fprintln(wx,DFS(buildTime, buildRule, winBuild))
	}
}
func DFS(buildTime []int, buildRule [][]int, winBuild int)int{
	stack := make([]int,0)
	stack = append(stack,winBuild)
	check := make([]int,len(buildTime)+1)
	answer := 0
	copy(check,buildTime)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range buildRule[node]{
			if check[neighbor] < check[node] + buildTime[neighbor]{
				check[neighbor] = check[node] + buildTime[neighbor]
			}
			if answer < check[neighbor]{
				answer = check[neighbor]
			}
			stack=append(stack,neighbor)
		}
	}
	return answer
}

func TS(buildTime []int, buildRule [][]int, inDegree []int, winBuild int)int{
	queue := make([]int,0)
	for j:=0 ; j<len(buildTime) ;j++{
		if inDegree[j] == 0 {
			queue = append(queue,j)
		}
	}
	rst := make([]int,len(buildTime))
	copy(rst,buildTime)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range buildRule[node]{
			inDegree[neighbor]--
			if rst[neighbor] < rst[node]+buildTime[neighbor]{
				rst[neighbor] = rst[node]+buildTime[neighbor]
			}
			if inDegree[neighbor] == 0 {
				queue=append(queue,neighbor)
			}
		}
	}
	return rst[winBuild]
}