package piscine

import (
	"ft"
)

func PrintAlphabet() {
	for a := 'a'; a <= 'z'; a++ {
		ft.PrintRune(a)
	}
	ft.PrintRune('\n')
}
