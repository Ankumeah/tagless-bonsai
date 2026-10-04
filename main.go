package main

import (
	"syscall/js"
)

var width uint
var height uint

var intervalID js.Value
var intervalFunc js.Func
var intervalFuncActive bool

func stopGrowth() {
	if intervalID.Truthy() {
		js.Global().Call("clearInterval", intervalID)
		intervalID = js.Undefined()
	}

	if intervalFuncActive {
		intervalFunc.Release()
		intervalFuncActive = false
	}
}

func startGrowth(output js.Value) {
	stopGrowth()

	if height < 7 || width < 7 {
		output.Set("textContent", "Screen width is too small")
		return
	}

	treeHeight := int(height - 3)
	treeWidth := int(width)

	bonsai := make([][]byte, treeHeight)

	for y := range bonsai {
		bonsai[y] = make([]byte, treeWidth)

		for x := range bonsai[y] {
			bonsai[y][x] = ' '
		}
	}

	root := node{
		x:  treeWidth / 2,
		y:  treeHeight - 1,
		dx: 0,
		dy: -1,
	}
	openNodes := []node{root}

	bonsai[root.y][root.x] = '|'

	intervalFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		text, nextNodes, done := genBonsai(bonsai, openNodes)

		output.Set("textContent", text)
		openNodes = nextNodes

		if done {
			stopGrowth()
		}

		return nil
	})

	intervalFuncActive = true

	intervalID = js.Global().Call(
		"setInterval",
		intervalFunc,
		1000,
	)
}

func setup(this js.Value, args []js.Value) any {
	document := js.Global().Get("document")
	body := document.Get("body")

	favicon := document.Call("createElement", "link")
	favicon.Set("rel", "icon")
	favicon.Set("type", "image/x-icon")
	favicon.Set("href", "/favicon.ico")

	title := document.Call("createElement", "h1")
	title.Set("textContent", "Tagless* Bonsai")
	title.Get("style").Set("color", "white")
	title.Get("style").Set("textAlign", "center")

	subtitle := document.Call("createElement", "h2")
	subtitle.Set("textContent", "The world's ugliest, yet the best bonsai")
	subtitle.Get("style").Set("color", "white")
	subtitle.Get("style").Set("textAlign", "center")

	reset := document.Call("createElement", "button")
	reset.Set("textContent", "Start / Reset")
	reset.Set("className", "reset")
	reset.Get("style").Set("display", "block")
	reset.Get("style").Set("margin", "5% auto")
	reset.Get("style").Set("backgroundColor", "grey")
	reset.Get("style").Set("fontSize", "2rem")

	output := document.Call("createElement", "pre")
	output.Set("id", "bonsai")
	output.Set("className", "bonsai")
	output.Get("style").Set("fontFamily", "monospace")
	output.Get("style").Set("color", "white")
	output.Get("style").Set("fontSize", "32px")
	output.Get("style").Set("lineHeight", "1.2")
	output.Get("style").Set("textAlign", "left")
	output.Get("style").Set("width", "720px")
	output.Get("style").Set("height", "720px")
	output.Get("style").Set("margin", "0 auto")
	output.Get("style").Set("border", "1px solid white")
	output.Get("style").Set("whiteSpace", "pre")

	body.Call("append", title, subtitle, reset, output)
  document.Get("head").Call("appendChild", favicon)

	height, width = getDimensions(output)

	reset.Call(
		"addEventListener",
		"click",
		js.FuncOf(func(this js.Value, args []js.Value) any {
			output.Set("textContent", getPot(height, width))
			startGrowth(output)
			return nil
		}),
	)

	output.Set("textContent", getPot(height, width))

	return nil
}

func main() {
	setupFunc := js.FuncOf(setup)
	defer setupFunc.Release()

	js.Global().Set("setup", setupFunc)

	select {}
}
