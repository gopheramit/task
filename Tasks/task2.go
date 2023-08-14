package tasks

func Solution2(s string, a []string) int {
	count := 0
	eleCount := make(map[rune]int)
	for _, ele := range s {
		eleCount[ele] += 1
	}
	for _, recipe := range a {
		tempEleCount := make(map[rune]int)
		for key, value := range eleCount {
			tempEleCount[key] = value
		}
		canMake := true
		for _, ingrident := range recipe {
			if tempEleCount[ingrident] == 0 {
				canMake = false
				break

			}
			tempEleCount[ingrident] -= 1

		}

		if canMake {
			count += 1
		}
	}
	return count
}
