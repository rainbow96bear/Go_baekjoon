package main

import "testing"

func BenchmarkTestDFS(b *testing.B){
	for i:=0 ; i<b.N ; i++{
		DFS([]int{0,10,1,100,10},[][]int{{},{},{1},{1},{2,3}},4)
		DFS([]int{0,10,20,1,5,8,7,1,43},[][]int{{},{},{},{},{},{},{},{},{}},7)
		DFS([]int{0,1,2,3},[][]int{{},{},{1},{2}},1)
		DFS([]int{0,5,5,5,5},[][]int{{},{},{1},{1,2},{}},4)
		DFS([]int{0,100000,99999,99997,99994,99990},[][]int{{},{},{1},{1,2},{1,2,3},{1,2,3,4}},4)
	}
}

func BenchmarkTestTS(b *testing.B){
	for i:=0 ; i<b.N ; i++{
		TS([]int{0,10,1,100,10},[][]int{{},{2,3},{4},{4},{}},[]int{0,0,1,1,2},4)
		TS([]int{0,10,20,1,5,8,7,1,43},[][]int{{},{2,3},{4,5},{6},{},{7},{7},{8},{}},[]int{0,0,1,1,1,1,1,2,1},7)
		TS([]int{0,1,2,3},[][]int{{},{2},{3},{}},[]int{0,0,1,1},1)
		TS([]int{0,5,5,5,5},[][]int{{},{2,3},{3},{},{}},[]int{0,0,1,2,0},4)
		TS([]int{0,100000,99999,9997,9994,9990},[][]int{{},{2,3,4,5},{3,4,5},{4,5},{5},{}},[]int{0,0,1,2,3,4},4)
	}
}