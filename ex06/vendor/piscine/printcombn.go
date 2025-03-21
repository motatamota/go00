package piscine

import (
	"ft"
)

func Recursive(nb int, pos int, arr []int) {
	if nb == 1 {
		if pos == 0 {
			for arr[pos] = 0; arr[pos] <= 9-nb; arr[pos]++ {
				for n := 0; n <= pos; n++ {
					ft.PrintRune(rune(arr[n] + '0'))
				}
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}

		} else {
			for arr[pos] = arr[pos-1] + 1; arr[pos] <= 9; arr[pos]++ {
				for n := 0; n <= pos; n++ {
					ft.PrintRune(rune(arr[n] + '0'))
				}
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
		return
	}

	if pos == 0 {
		for arr[pos] = 0; arr[pos] <= 9-nb; arr[pos]++ {
			Recursive(nb-1, pos+1, arr)
		}
	} else {
		for arr[pos] = arr[pos-1] + 1; arr[pos] <= 9; arr[pos]++ {
			Recursive(nb-1, pos+1, arr)
		}
	}
}

func PrintCombN(nb int) {
	arr := make([]int, nb)

	Recursive(nb, 0, arr)
	for n := 0; n < nb; n++ {
		ft.PrintRune(rune(10 - nb + n + '0'))
	}
	ft.PrintRune('\n')
}
