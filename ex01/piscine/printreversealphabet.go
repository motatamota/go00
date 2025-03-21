package piscine

import (
	"ex01/ft"
)

func PrintReverseAlphabet() {
	for a := 'z'; a >= 'a'; a-- {
		ft.PrintRune(a)
	}
	ft.PrintRune('\n')
}
