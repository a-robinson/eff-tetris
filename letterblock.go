package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/util"
)

type letterBlock struct {
	letter rune
	color  eff.Color
	rect   eff.Rect
	mover  func(eff.Rect, float64) eff.Rect
}

func (l *letterBlock) draw(c eff.Canvas) {
	c.FillRect(l.rect, l.color)

	t := string(l.letter)
	lp, err := util.CenterTextInRect(t, l.rect, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	textColor := eff.Black()
	if rand.Intn(10) >= 5 {
		textColor = eff.White()
	}

	err = c.DrawText(t, textColor, lp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func createMover(start eff.Point, end eff.Point) func(eff.Rect, float64) eff.Rect {
	return func(r eff.Rect, p float64) eff.Rect {
		r.X = start.X + int(float64(end.X-start.X)*p)
		r.Y = start.Y + int(float64(end.Y-start.Y)*p)
		return r
	}
}

func letterBlocksForString(s string, offset eff.Point, c eff.Canvas, bottom bool) []letterBlock {
	var colors []eff.Color
	colors = append(colors, eff.Color{R: 45, G: 255, B: 254, A: 255})
	colors = append(colors, eff.Color{R: 11, G: 36, B: 251, A: 255})
	colors = append(colors, eff.Color{R: 253, G: 164, B: 40, A: 255})
	colors = append(colors, eff.Color{R: 255, G: 253, B: 56, A: 255})
	colors = append(colors, eff.Color{R: 41, G: 253, B: 47, A: 255})
	colors = append(colors, eff.Color{R: 169, G: 38, B: 251, A: 255})
	colors = append(colors, eff.Color{R: 252, G: 13, B: 27, A: 255})

	blockSize := 30

	var letterBlocks []letterBlock
	for i, letter := range s {
		start := eff.Point{
			X: offset.X + i*blockSize,
			Y: -(i * i * blockSize) - blockSize,
		}

		if bottom {
			start.Y += c.Height()
		}

		end := eff.Point{
			X: offset.X + i*blockSize,
			Y: offset.Y,
		}

		letterBlocks = append(letterBlocks, letterBlock{
			letter: letter,
			color:  colors[rand.Intn(len(colors))],
			rect: eff.Rect{
				X: start.X,
				Y: start.Y,
				W: blockSize,
				H: blockSize,
			},
			mover: createMover(start, end),
		})
	}

	return letterBlocks
}
