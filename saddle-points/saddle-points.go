// Package matrix
package matrix

type Pair struct {
	x, y int // a point in the matrix
}

// Saddle gets the saddle point in a matrix
func (m Matrix) Saddle() []Pair {

	var sp []Pair
	rows := m.Rows()
	cols := m.Cols()

	for i, row := range rows {
		bigI, bigV := biggest(row)
		if smallest(cols[bigI], bigV) {
			sp = append(sp, Pair{i, bigI})
		}
	}
	return sp
}

// biggest return the index and value
// of the largest number in a matrix row
func biggest(x []int) (int, int) {
	var n, bigI, bigV int
	for i, v := range x {
		if v > n {
			n = v
			bigI = i
			bigV = n
		}
	}
	return bigI, bigV
}

// smallest checks if a given number is the smallest
// in a matrix column
func smallest(x []int, rowV int) bool {

	var isSmallest bool

	for _, v := range x {
		if v < rowV {
			isSmallest = false
			return isSmallest
		} else {
			isSmallest = true
		}
	}
	return isSmallest
}
