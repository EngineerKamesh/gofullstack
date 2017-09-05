package main

import "honnef.co/go/js/dom"

func toggleIsomorphicGopher(isVisible bool) {

	d := dom.GetWindow().Document()
	isomorphicGopherImage := d.GetElementByID("isomorphicGopher").(*dom.HTMLImageElement)

	if isVisible == true {
		isomorphicGopherImage.Style().SetProperty("display", "inline", "")
	} else {
		isomorphicGopherImage.Style().SetProperty("display", "none", "")
	}

}

func ShowIsomorphicGopher() {
	toggleIsomorphicGopher(true)
}

func HideIsomorphicGopher() {
	toggleIsomorphicGopher(false)
}
