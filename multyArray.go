package main

import "fmt"

func main() {
	fmt.Println("다차원 배열")
	//일차원 배열

	// var singleArray = [3]int{1, 2, 3}
	// fmt.Println(singleArray)
	var singleArray = [100]int{99: 5}
	fmt.Println(singleArray)

	// 2차원 배열 실습
	var doubleArray = [2][3]int{
		{1: 55, 2: 6},
		{4, 5, 6},
	}
	fmt.Println(doubleArray)

	// 차원 배열 실습

	var multyArray = [2][3][2]int{
		{{1, 2}, {3, 4}, {5, 6}},
		{{7, 8}, {9, 10}, {11, 12}},
	}
	fmt.Println(multyArray)
}
