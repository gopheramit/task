package main

import (
	"fmt"

	tasks "github.com/gopheramit/task/Tasks"
)

func main() {
	//.fmt.Println("hello world")
	// A := []int{100, 100, 100, -10}
	// D := []string{"2020-12-31", "2020-12-22", "2020-12-03", "2020-12-29"}
	// A := []int{180, -50, -25, -25}
	// D := []string{"2020-01-01", "2020-01-01", "2020-01-01", "2020-01-31"}
	// // A := []int{1, -1, 0, -105, 1}
	// // D := []string{"2020-12-31", "2020-04-04", "2020-04-04", "2020-04-14", "2020-07-12"}
	// // A := []int{100, 100, -10, -20, -30}
	// // D := []string{"2020-01-01", "2020-02-01", "2020-02-11", "2020-02-05", "2020-02-08"}
	// // A := []int{-60, 60, -40, -20}
	// // D := []string{"2020-10-01", "2020-02-02", "2020-10-10", "2020-10-30"}
	// fmt.Println(tasks.Solution(A, D))
	fmt.Println("enter the print the pyramid")
	var num int
	fmt.Scanf("%d", &num)
	tasks.PyramidSolution(num)

}
