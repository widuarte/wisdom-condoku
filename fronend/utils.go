package fronend

import (
	"wisdom_condoku/console"
	"wisdom_condoku/model"
)

func PrintCell(cell *model.Cell) {
	getColor(cell)(getValue(cell))
}

func getValue(cell *model.Cell) string {
	if cell.Value == "" {
		return "_"
	}
	return cell.Value
}

func getColor(cell *model.Cell) console.Color {
	if cell.Value == "" {
		return console.Blue
	}
	switch cell.State {
	case model.NEWER:
		return console.Green
	case model.ERROR:
		return console.Red
	case model.DEFAULT:
		return console.White
	}
	return nil
}
