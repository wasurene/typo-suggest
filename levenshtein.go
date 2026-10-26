package main

func min(a, b, c int) int {
	min := a

	slice := []int{a, b, c}

	for i := range slice {
		if min > slice[i] {
			min = slice[i]
		}
	}

	return min
}

func levenshtein(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)
	N := len(runesA)
	M := len(runesB)

	matrix := make([][]int, N+1)

	for i := range matrix {
		matrix[i] = make([]int, M+1)
	}

	for j := 0; j <= M; j++ {
		matrix[0][j] = j
	}

	for i := 0; i <= N; i++ {
		matrix[i][0] = i
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if runesA[i-1] == runesB[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1]) + 1
			}
		}
	}

	return matrix[N][M]

}
