package main // OMIT

import "net/http"

func MyExportedFunc1(url string) {
	r, _ := http.Get(url)          // blocking call
	println(r)
}
