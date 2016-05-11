package main // OMIT

import "net/http"

func MyExportedFunc2(url string) {
	go func() {
		r, _ := http.Get(url)      // blocking call
		println(r)
	}()
}
