package main

import (
	"math/rand"
	"strings"
	"syscall/js"
)

type node struct {
	x  int
	y  int
	dx int
	dy int
}

var symbols = map[int]byte{
	-1: '/',
	0:  '|',
	1:  '\\',
}

func getDimensions(output js.Value) (height uint, width uint) {
	const fontSize = 32.0
	const lineHeight = fontSize * 1.2
	const charWidth = fontSize * 0.6

	paragraphHeight := output.Get("clientHeight").Float()
	paragraphWidth := output.Get("clientWidth").Float()

	height = uint(paragraphHeight / lineHeight)
	width = uint(paragraphWidth / charWidth)

	return height, width
}

func getPot(height uint, width uint) string {
	if height < 7 || width < 7 {
		return "Screen width is too small"
	}

	var pot strings.Builder

	for range height - 3 {
		pot.WriteByte('\n')
	}

	pot.WriteByte(' ')

	for range width - 2 {
		pot.WriteByte('-')
	}

	pot.WriteByte(' ')
	pot.WriteByte('\n')

	pot.WriteByte(' ')
	pot.WriteByte('\\')

	for range width - 4 {
		pot.WriteByte(' ')
	}

	pot.WriteByte('/')
	pot.WriteByte(' ')
	pot.WriteByte('\n')

	pot.WriteByte(' ')
	pot.WriteByte(' ')

	for range width - 4 {
		pot.WriteByte('-')
	}

	pot.WriteByte(' ')
	pot.WriteByte(' ')

	return pot.String()
}

func genBonsai(bonsai [][]byte, openNodes []node) (string, []node, bool) {
	if len(openNodes) == 0 {
		return renderBonsai(bonsai), nil, true
	}

	nodes := make([]node, 0, len(openNodes)+2)

	for _, current := range openNodes {
		n := current

		// Change direction
		switch rand.Intn(3) {
		case 0:
			n.dx = -1
		case 1:
			n.dx = 0
		case 2:
			n.dx = 1
		}

		n.x += n.dx
		n.y += n.dy

		// Close this branch if it leaves the area
		if n.x < 0 ||
			n.x >= len(bonsai[0]) ||
			n.y < 0 ||
			n.y >= len(bonsai) {
			continue
		}

		// Close a branch
		if rand.Intn(9) > 8 {
			continue
		}

		bonsai[n.y][n.x] = symbols[n.dx]
		nodes = append(nodes, n)

		// Split branch
		if rand.Intn(2) == 0 {
			branchDirection := -1

			if n.dx < 0 {
				branchDirection = 1
			}

			nodes = append(nodes, node{
				x:  n.x,
				y:  n.y,
				dx: branchDirection,
				dy: -1,
			})
		}
	}

	return renderBonsai(bonsai), nodes, len(nodes) == 0
}

func renderBonsai(bonsai [][]byte) string {
	if len(bonsai) == 0 || len(bonsai[0]) == 0 {
		return ""
	}

	var output strings.Builder

	for _, row := range bonsai {
		output.Write(row)
		output.WriteByte('\n')
	}

	pot := getPot(
		uint(len(bonsai))+3,
		uint(len(bonsai[0])),
	)

	pot = strings.TrimLeft(pot, "\n")

	output.WriteString(pot)

	return output.String()
}
