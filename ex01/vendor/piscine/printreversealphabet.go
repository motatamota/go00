package piscine

import (
	"ft"
)

func PrintReverseAlphabet() {
	for a := 'z'; a >= 'a'; a-- {
		ft.PrintRune(a)
	}
	ft.PrintRune('\n')
}
