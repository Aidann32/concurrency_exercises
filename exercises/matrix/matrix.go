package matrix

import "errors"

func checkMultiplicability(a [][]int, b [][]int) (bool, error) {
	bRows := len(b)
	aCols := len(a[0])
	bCols := len(b[0])
	for _, arr := range a {
		if len(arr) != aCols {
			return false, errors.New("a is not matrix")
		}
	}
	for _, arr := range b {
		if len(arr) != bCols {
			return false, errors.New("b is not matrix")
		}
	}
	if aCols != bRows {
		return false, nil
	}
	return true, nil
}

func calculateResultElement(a [][]int, b [][]int, resRowIndex, resColIndex int) int {
	result := 0
	for i := 0; i < len(a[0]); i++ {
		result += a[resRowIndex][i] * b[i][resColIndex]
	}
	return result
}

func multiplyMatrices(a [][]int, b [][]int, p int) ([][]int, error) {
	isMultiplicable, err := checkMultiplicability(a, b)
	if !isMultiplicable {
		return nil, errors.New("matrices are not multiplicable")
	}
	if err != nil {
		return nil, err
	}

	cRows, cCols := len(a), len(b[0])
	resultAccess := make([][]bool, cRows)
	result := make([][]int, cRows)
	for i := range resultAccess {
		resultAccess[i] = make([]bool, cCols)
		result[i] = make([]int, 0)
	}

	// Worker pool of p goroutines which check accessMatrix and calculate c[i][j] if accessMatrix[i][j] is false
	// true means that element c[i][j] is calculated or calculating

	return result, nil
}

func Run(goroutineNumber int, a [][]int, b [][]int) ([][]int, error) {
	result, err := multiplyMatrices(a, b, goroutineNumber)
	if err != nil {
		return nil, err
	}
	return result, nil
}
