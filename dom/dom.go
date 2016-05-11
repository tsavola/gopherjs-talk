package main

import "github.com/gopherjs/gopherjs/js"

var doc = js.Global.Get("document")

func main() {
	div := doc.Call("createElement", "div")
	div.Set("innerText", "hello world")

	doc.Get("body").Call("appendChild", div)
}
