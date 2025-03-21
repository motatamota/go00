package piscine

import (
	"ex05/ft"
)

func PrintComb2() {
	x := 0
	y := 0
	for x < 98{
		y = x + 1
		for y < 100{
			ft.PrintRune(rune(x / 10 + '0'))
			ft.PrintRune(rune(x % 10 + '0'))
			ft.PrintRune(' ')
			ft.PrintRune(rune(y / 10 + '0'))
			ft.PrintRune(rune(y % 10 + '0'))
			ft.PrintRune(',')
			ft.PrintRune(' ')
			y++
		}
		x++
	}
	ft.PrintRune(rune(x / 10 + '0'))
	ft.PrintRune(rune(x % 10 + '0'))
	ft.PrintRune(' ')
	ft.PrintRune(rune((y - 1) / 10 + '0'))
	ft.PrintRune(rune((y - 1) % 10 + '0'))
	ft.PrintRune('\n')
}
