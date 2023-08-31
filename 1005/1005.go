package main

import (
	"bufio"
	"fmt"
	"os"
)

type stackArr struct{
	value []int
}
func (s *stackArr)pop()int{
	rst := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return rst
}
func (s *stackArr)push(v int){
	s.value = append(s.value,v)
}

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
		buildTime := make([]int,buildNum)
		buildRule := make(map[int][]int,ruleNum)
		for j := 0 ; j<buildNum ; j++{
			fmt.Fscan(rx,&buildTime[j])
		}
		for k:=0 ; k<ruleNum ; k++{
			var key, value int
			fmt.Fscan(rx, &value, &key)
			buildRule[key]=append(buildRule[key],value)
		}
		fmt.Fscan(rx,&winBuild)
		fmt.Fprintln(wx,DFS(buildRule,buildTime,winBuild))
	}
}

func DFS(buildRule map[int][]int, buildTime []int,winBuild int)int{
	var stack stackArr
	check := []int{}
	rst := 0
	answer := 0
	stack.push(winBuild)
	for len(stack.value) != 0 {
		var now = stack.pop()
		rst = check[now] + buildTime[now-1]
		if buildRule[now] == nil{
			if answer < rst{
				answer = rst
			}
		}else{
			for i:=0 ; i<len(buildRule[now]) ; i++{
				check[buildRule[now][i]] = rst
				stack.push(buildRule[now][i])
			}
		}
	}
	return answer
}