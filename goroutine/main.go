package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("window").Set("MyStuff", map[string]interface{}{
		"f1": MyExportedFunc1,
		"f2": MyExportedFunc2,
	})
}
