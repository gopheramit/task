package tasks

import "fmt"

func Solution(A []int, D []string) int {
	bal := 0
	deduction := 0
	noOfTransactions := map[string]int{}
	totalAmount := map[string]int{}

	for i, _ := range A {
		if A[i] >= 0 {
			bal += A[i]
		} else {
			bal += A[i]
			date := D[i]
			month := date[5:7]
			noOfTransactions[month] += 1
			totalAmount[month] -= A[i]
		}
	}
	fmt.Println(totalAmount)
	for k, v := range noOfTransactions {
		if v < 3 {
			deduction += 5
		} else if v >= 3 && totalAmount[k] < 100 {
			deduction += 5
		}
	}
	monthWithoutCardPayments := 12 - len(noOfTransactions)
	deduction += 5 * monthWithoutCardPayments
	bal = bal - deduction
	return bal

}
