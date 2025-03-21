package piscine

import (
	"ex02/ft"
)

func PrintReverseAlphabet() {
	for a := 0; a <= 9; a++ {
		ft.PrintRune(rune(a + '0'))
	}
	ft.PrintRune('\n')
}
