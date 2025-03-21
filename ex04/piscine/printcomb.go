package piscine

import (
	"ex04/ft"
)

func PrintComb() {
	x := 0
	y := 0
	z := 0
	for x < 7 {
		y = x + 1
		for y <= 8 {
			z = y + 1
			for z <= 9 {
				ft.PrintRune(rune(x + '0'))
				ft.PrintRune(rune(y + '0'))
				ft.PrintRune(rune(z + '0'))
				ft.PrintRune(',')
				ft.PrintRune(' ')
				z++
			}
			y++
		}
		x++
	}
	ft.PrintRune(rune(x + '0'))
	ft.PrintRune(rune(y + '0' - 1))
	ft.PrintRune(rune(z + '0' - 1))
	ft.PrintRune('\n')
}
