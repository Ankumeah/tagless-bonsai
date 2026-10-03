package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func getDimensions(output js.Value) (height uint, width uint) {
	const fontSize = 32.0
	const lineHeight = fontSize * 1.2
	const charWidth = fontSize * 0.6

	paragraphHeight := output.Get("clientHeight").Float()
	paragraphWidth := output.Get("clientWidth").Float()

	height = uint(paragraphHeight / lineHeight)
	width = uint(paragraphWidth / charWidth)

	return
}

func getPot(height uint, width uint) string {
	if height < 7 || width < 7 {
		return "Screen width too small"
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

	potString := pot.String()
	fmt.Printf("Pot:\n%v\n", potString)

	return potString
}
