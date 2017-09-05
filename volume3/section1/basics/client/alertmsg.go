package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func DisplayAlertMessageJSGlobal(message string) {
	js.Global.Call("alert", message)
}

func DisplayAlertMessageDOM(message string) {
	dom.GetWindow().Alert(message)
}
