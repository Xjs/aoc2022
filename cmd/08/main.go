package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Xjs/aoc2021/integer"
	"github.com/Xjs/aoc2021/integer/grid"
	"github.com/Xjs/aoc2021/part"
)

func main() {
	g, err := integer.ReadGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	w := g.Width()
	h := g.Height()

	if !part.One() {
		var scenicP grid.Point
		var scenicScore int

		for i := uint(0); i < w; i++ {
			for j := uint(0); j < h; j++ {
				x := i
				y := j
				p := grid.P(x, y)
				treeHeight := g.MustAt(p)

				left := 0
				// left
				for xn := int(x) - 1; xn >= 0; xn-- {
					left++
					if newHeight := g.MustAt(grid.P(uint(xn), y)); newHeight >= treeHeight {
						break
					}
				}

				right := 0
				// right
				for xn := x + 1; xn < w; xn++ {
					right++
					if newHeight := g.MustAt(grid.P(uint(xn), y)); newHeight >= treeHeight {
						break
					}

				}

				up := 0
				// up
				for yn := int(y) - 1; yn >= 0; yn-- {
					up++
					if newHeight := g.MustAt(grid.P(x, uint(yn))); newHeight >= treeHeight {
						break
					}
				}

				down := 0
				// down
				for yn := y + 1; yn < h; yn++ {
					down++
					if newHeight := g.MustAt(grid.P(x, uint(yn))); newHeight >= treeHeight {
						break
					}
				}

				score := left * right * up * down
				if score > scenicScore {
					scenicScore = score
					scenicP = p
				}
			}
		}

		fmt.Println(scenicP, scenicScore)
		return
	}

	visible := make(map[grid.Point]struct{})

	// top to bottom
	for i := uint(0); i < w; i++ {
		highest := -1
		for j := uint(0); j < h; j++ {
			x := i
			y := j
			p := grid.P(x, y)

			if h := g.MustAt(p); h > highest {
				highest = h
				visible[p] = struct{}{}
			}
		}
	}

	// bottom to top
	for i := uint(0); i < w; i++ {
		highest := -1
		for j := uint(0); j < h; j++ {
			x := w - i - 1
			y := h - j - 1
			p := grid.P(x, y)

			if h := g.MustAt(p); h > highest {
				highest = h
				visible[p] = struct{}{}
			}
		}
	}

	// left to right
	for j := uint(0); j < h; j++ {
		highest := -1
		for i := uint(0); i < w; i++ {
			x := i
			y := j
			p := grid.P(x, y)

			if h := g.MustAt(p); h > highest {
				highest = h
				visible[p] = struct{}{}
			}
		}
	}

	// right to left
	for j := uint(0); j < h; j++ {
		highest := -1
		for i := uint(0); i < w; i++ {
			x := w - i - 1
			y := h - j - 1
			p := grid.P(x, y)

			if h := g.MustAt(p); h > highest {
				highest = h
				visible[p] = struct{}{}
			}
		}
	}

	fmt.Println(len(visible))
}
