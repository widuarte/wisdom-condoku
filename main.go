package main

import (
	"wisdom_condoku/fronend"
	"wisdom_condoku/model"
)

func main() {
	matrix := model.NewMatrix()
	fronend.PrintBoard(matrix)
}
