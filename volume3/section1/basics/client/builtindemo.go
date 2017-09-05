package main

import (
	"github.com/gopherjs/jsbuiltin"
	"honnef.co/go/js/dom"
)

func builtinDemo(element dom.Element) {

	if jsbuiltin.TypeOf(element) == "object" {
		println("Using the typeof operator, we can see that the element that was clicked, is an object.")
	}

}
