// Package matrix
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

type Matrix [][]int

// New creates a matrix
func New(s string) (*Matrix, error) {
	mys := strings.Split(s, "\n")
	m := make(Matrix, len(mys))

	for i := range mys {
		srow := strings.Split(strings.TrimSpace(mys[i]), " ")
		irow := make([]int, len(srow))
		//	fmt.Println(len(srow), srow, irow)

		for j := range srow {
			integer, err := strconv.Atoi(srow[j])
			if err != nil {
				return &m, err
			}
			irow[j] = integer
			//	fmt.Println(integer, irow)
		}
		if i != 0 {
			if len(irow) != len(m[i-1]) {
				errUneven := errors.New("uneven rows")
				return &m, errUneven
			}
		}
		m[i] = irow

	}

	return &m, nil
}

// Rows copies the values in a matrix ordered by rows
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		rows[i] = make([]int, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}

	return rows
}

// Cols copies the values in a matrix ordered by colums
func (m Matrix) Cols() [][]int {

	cols := make([][]int, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(m); j++ {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

// Set sets the value at a specific point in a matrix
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true

}
