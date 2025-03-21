package piscine

import (
	"00_01/main/ft"
)

func PrintReverseAlphabet() {
	for a := 'z'; a >= 'a'; a-- {
		ft.PrintRune(a)
	}
	ft.PrintRune('\n')
}
