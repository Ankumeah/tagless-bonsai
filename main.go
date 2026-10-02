package main

import (
	"syscall/js"
)

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

	canvas := document.Call("createElement", "canvas")
	canvas.Set("id", "bonsai")
	canvas.Set("width", 720)
	canvas.Set("height", 720)
	canvas.Get("style").Set("display", "block")
	canvas.Get("style").Set("margin", "0 auto")
	canvas.Get("style").Set("border", "1px solid white")

	context := canvas.Call("getContext", "2d")
	context.Set("fillStyle", "white")
	context.Set("font", "32px Arial")

	reset.Call(
		"addEventListener", "click",
		js.FuncOf(
			func(this js.Value, args []js.Value) any {
				context.Call(
					"clearRect",
					0, 0,
					canvas.Get("width"), canvas.Get("height"),
				)
				context.Call("fillText", "Hello world!", 40, 100) // TODO: remove
				return nil
			},
		),
	)

	body.Call("append", reset, canvas)

	return nil
}

func main() {
	setupFunc := js.FuncOf(setup)
	defer setupFunc.Release()

	js.Global().Set("setup", setupFunc)

	select {}
}
