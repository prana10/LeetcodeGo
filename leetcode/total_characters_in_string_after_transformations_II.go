package leetcode

const MOD = 1000000007

func LengthAfterTransformations(s string, t int, nums []int) int {
	totalLen := 0
	charCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	matrix := make([][]int, 26)
	for i := 0; i < 26; i++ {
		matrix[i] = make([]int, 26)
		numChars := nums[i]
		for j := 0; j < numChars; j++ {
			nextChar := (i + j + 1) % 26
			matrix[i][nextChar]++
		}
	}

	vector := make([]int, 26)
	for c, count := range charCount {
		vector[c-'a'] = count
	}

	resultMatrix := matrixPower(matrix, int64(t))
	resultVector := make([]int, 26)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			resultVector[i] = (resultVector[i] + ((vector[j]%MOD)*(resultMatrix[j][i]%MOD))%MOD) % MOD
		}
	}

	for i := 0; i < 26; i++ {
		totalLen = (totalLen + resultVector[i]) % MOD
	}

	return totalLen
}

func multiplyMatrix(a, b [][]int) [][]int {
	n := len(a)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				res[i][j] = (res[i][j] + ((a[i][k]%MOD)*(b[k][j]%MOD))%MOD) % MOD
			}
		}
	}

	return res
}

func matrixPower(matrix [][]int, n int64) [][]int {
	size := len(matrix)
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
		result[i][i] = 1
	}

	for n > 0 {
		if n%2 == 1 {
			result = multiplyMatrix(result, matrix)
		}
		matrix = multiplyMatrix(matrix, matrix)
		n /= 2
	}

	return result
}
