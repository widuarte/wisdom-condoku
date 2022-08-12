package model

type Matrix [9][9]*Cell

func NewMatrix() *Matrix {
	m := &Matrix{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			m[i][j] = NewCell("", true)
		}
	}
	return m
}
