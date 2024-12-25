package matrix

import (
	"errors"
	"fmt"
	"sync"
)

type matrixIndex struct {
	I int
	J int
}

type Multiplication struct {
	A            [][]int
	B            [][]int
	P            int
	resultMatrix [][]int
	resultAccess [][]bool

	wg    *sync.WaitGroup
	mutex *sync.RWMutex
	jobs  chan matrixIndex
}

func (m *Multiplication) checkMultiplicability() (bool, error) {
	bRows := len(m.B)
	aCols := len(m.A[0])
	bCols := len(m.B[0])
	for _, arr := range m.A {
		if len(arr) != aCols {
			return false, errors.New("a is not matrix")
		}
	}
	for _, arr := range m.B {
		if len(arr) != bCols {
			return false, errors.New("b is not matrix")
		}
	}
	if aCols != bRows {
		return false, nil
	}
	return true, nil
}

func (m *Multiplication) calculateResultElement(resRowIndex, resColIndex int) int {
	result := 0
	for i := 0; i < len(m.A[0]); i++ {
		result += m.A[resRowIndex][i] * m.B[i][resColIndex]
		fmt.Printf("Calculating A[%d][%d](%d) and B[%d][%d](%d) result is %d\n",
			resRowIndex, i, m.A[resRowIndex][i], i, resColIndex, m.B[i][resColIndex], result)
	}
	return result
}

func (m *Multiplication) getAccessMatrixElement(i, j int) bool {
	defer m.mutex.RUnlock()
	m.mutex.RLock()
	return m.resultAccess[i][j]
}

func (m *Multiplication) setAccessMatrixElement(i, j int, value bool) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	m.resultAccess[i][j] = value
}

func (m *Multiplication) worker() {
	for j := range m.jobs {
		defer m.wg.Done()
		indexes := j
		if !m.getAccessMatrixElement(indexes.I, indexes.J) {
			m.setAccessMatrixElement(indexes.I, indexes.J, true)
			result := m.calculateResultElement(indexes.I, indexes.J)
			m.resultMatrix[indexes.I][indexes.J] = result
		}
	}
}

func (m *Multiplication) multiplyMatrices() {

	// Worker pool of p goroutines which check accessMatrix and calculate c[i][j] if accessMatrix[i][j] is false
	// true means that element c[i][j] is calculated or calculating

	m.wg.Add(len(m.resultMatrix) * len(m.resultMatrix[0]))

	for i := 1; i <= m.P; i++ {
		go m.worker()
	}

	for i := 0; i < len(m.resultMatrix); i++ {
		for j := 0; j < len(m.resultMatrix[0]); j++ {
			m.jobs <- matrixIndex{I: i, J: j}
		}
	}
	close(m.jobs)
	m.wg.Wait()
}

func NewMultiplication(a, b [][]int, p int) *Multiplication {
	return &Multiplication{
		A: a,
		B: b,
		P: p,
	}
}

func (m *Multiplication) Run() ([][]int, error) {
	isMultiplicable, err := m.checkMultiplicability()
	if !isMultiplicable {
		return nil, errors.New("matrices are not multiplicable")
	}
	if err != nil {
		return nil, err
	}

	cRows, cCols := len(m.A), len(m.B[0])
	m.resultAccess = make([][]bool, cRows)
	m.resultMatrix = make([][]int, cRows)
	for i := range m.resultAccess {
		m.resultAccess[i] = make([]bool, cCols)
		m.resultMatrix[i] = make([]int, cCols)
	}

	m.jobs = make(chan matrixIndex, cRows*cCols)
	m.wg = &sync.WaitGroup{}
	m.mutex = &sync.RWMutex{}

	m.multiplyMatrices()
	return m.resultMatrix, nil
}
