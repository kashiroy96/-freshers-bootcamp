package main

import "fmt"

type matrix struct {
	rows  int
	cols  int
	point [][]int
}

func (m matrix) lenr() int {
	return m.rows
}

func (m matrix) lenc() int {
	return m.cols
}

func (m matrix) set(i, j, val int) {
	m.point[i][j] = val
}

func (m matrix) add(arr [][]int) matrix {

	var r = m.rows
	var c = m.cols
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m.point[i][j] = m.point[i][j] + arr[i][j]
		}
	}
	return m
}

func (m matrix) show() {
	var r = m.rows
	fmt.Printf("\n")
	for i := 0; i < r; i++ {
		fmt.Println(m.point[i])
	}
}

func main() {

	m := matrix{3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
	fmt.Println("number of rows: ", m.lenr())
	fmt.Println("number of columns: ", m.lenc())
	fmt.Println("\n")
	fmt.Println("set the values ")
	fmt.Println("old: ", m.point[1][2])
	m.set(1, 2, 90)
	fmt.Println("updated: ", m.point[1][2])
	fmt.Println("sum :")
	fmt.Println("\n")
	m.add([][]int{{11, 12, 13}, {14, 15, 16}, {17, 18, 19}})
	m.show()

}
