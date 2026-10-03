package main

import (
	"syscall/js"
)

var width uint
var height uint

func setup(this js.Value, args []js.Value) any {
	document := js.Global().Get("document")
	body := document.Get("body")

	reset := document.Call("createElement", "button")
	reset.Set("textContent", "Reset")
	reset.Set("className", "reset")
	reset.Get("style").Set("display", "block")
	reset.Get("style").Set("margin", "5% auto")
	reset.Get("style").Set("backgroundColor", "grey")
	reset.Get("style").Set("fontSize", "2rem")

	output := document.Call("createElement", "p")
	output.Set("id", "bonsai")
	output.Set("className", "bonsai")
	output.Get("style").Set("whiteSpace", "pre")
	output.Get("style").Set("fontFamily", "monospace")
	output.Get("style").Set("color", "white")
	output.Get("style").Set("fontSize", "32px")
	output.Get("style").Set("lineHeight", "1.2")
	output.Get("style").Set("textAlign", "left")
	output.Get("style").Set("width", "720px")
	output.Get("style").Set("height", "720px")
	output.Get("style").Set("margin", "0 auto")
	output.Get("style").Set("border", "1px solid white")

	reset.Call(
		"addEventListener", "click",
		js.FuncOf(
			func(this js.Value, args []js.Value) any {
				output.Set("textContent", getPot(height, width))
				return nil
			},
		),
	)

	body.Call("append", reset, output)

	height, width = getDimensions(output)

	return nil
}

func main() {
	setupFunc := js.FuncOf(setup)
	defer setupFunc.Release()

	js.Global().Set("setup", setupFunc)

	select {}
}
