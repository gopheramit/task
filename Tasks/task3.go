package tasks

import "fmt"

func PyramidSolution(n int) {
	// n := 10
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	// fmt.Println()
}
