package fronend

import (
	"wisdom_condoku/console"
	"wisdom_condoku/model"
)

func PrintTitle() {
	console.White("\n░█████╗░░█████╗░███╗░░██╗██████╗░░█████╗░██╗░░██╗██╗░░░██╗\n██╔══██╗██╔══██╗████╗░██║██╔══██╗██╔══██╗██║░██╔╝██║░░░██║\n██║░░╚═╝██║░░██║██╔██╗██║██║░░██║██║░░██║█████═╝░██║░░░██║\n██║░░██╗██║░░██║██║╚████║██║░░██║██║░░██║██╔═██╗░██║░░░██║\n╚█████╔╝╚█████╔╝██║░╚███║██████╔╝╚█████╔╝██║░╚██╗╚██████╔╝\n░╚════╝░░╚════╝░╚═╝░░╚══╝╚═════╝░░╚════╝░╚═╝░░╚═╝░╚═════╝░")
}

func PrintMatrix(matrix *model.Matrix) {
	for i := 0; i < 9; i++ {
		print("        ")
		if i%3 == 0 {
			console.Blue("____________  ___________  ____________")
			println()
			print("        ")
		}
		for j := 0; j < 9; j += 3 {
			console.Blue("| ")
			PrintCell(matrix[i][j])
			console.Blue(" | ")
			PrintCell(matrix[i][j+1])
			console.Blue(" | ")
			PrintCell(matrix[i][j+2])
			console.Blue(" |")
		}
		println()
	}
}

func PrintBoard(matrix *model.Matrix) {
	PrintTitle()
	println()
	println()
	PrintMatrix(matrix)
	println()
	println()
}
