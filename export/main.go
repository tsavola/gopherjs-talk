package main // OMIT

import "github.com/gopherjs/gopherjs/js" // OMIT

const MyString = "hello world"

func MyApplyFunc(callable *js.Object, arg string) {
	if callable != js.Undefined {
		callable.Invoke(arg)
	}
}

func main() {
	js.Global.Get("window").Set("MyStuff", map[string]interface{}{
		"MY_STRING":   MyString,
		"myApplyFunc": MyApplyFunc,
		"NewThing1":   NewThing1, // OMIT
		"NewThing2":   NewThing2, // OMIT
	})
}
