package main

import (
	"fmt"
	"math"
	"os"
)

type xBeforeY struct {
	appearance string
	flipCount  int
}

func main() {
	str := os.Args[1]

	ch := make(chan *xBeforeY)

	go flipit(ch, str)

	var minimum int = math.MaxInt
	var minString string

	for p := range ch {
		fmt.Printf("%q\t%d\n", p.appearance, p.flipCount)
		if p.flipCount < minimum {
			minimum = p.flipCount
			minString = p.appearance
		}
	}

	fmt.Printf("minimum flips required: %d   %q\n", minimum, minString)
}

func flipit(ch chan *xBeforeY, str string) {
	runes := []rune(str)
	fliprecurse(ch, runes, len(runes), 0, 0)
	close(ch)
}

func fliprecurse(ch chan *xBeforeY, str []rune, ln int, idx int, flipCount int) {
	if correctlyFlipped(str) {
		ch <- &xBeforeY{
			appearance: string(str),
			flipCount:  flipCount,
		}
		return
	}
	if idx >= ln {
		return
	}

	// Leave rune at idx alone
	fliprecurse(ch, str, ln, idx+1, flipCount)

	// flip rune at idx
	holder := str[idx]
	if str[idx] == 'x' {
		str[idx] = 'y'
	} else {
		str[idx] = 'x'
	}
	fliprecurse(ch, str, ln, idx+1, flipCount+1)
	str[idx] = holder
}

func correctlyFlipped(str []rune) bool {
	firstY := -1
	lastX := -1
	for n := range str {
		if str[n] == 'y' {
			firstY = n
		}
		if str[n] == 'x' {
			lastX = n
			if firstY >= 0 && lastX > firstY {
				return false
			}
		}
	}

	if firstY < 0 || lastX < 0 {
		return true
	}

	return lastX < firstY
}
