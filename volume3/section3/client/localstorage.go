package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var localStorage = js.Global.Get("localStorage")
var D = dom.GetWindow().Document().(dom.HTMLDocument)

func run() {
	saveButton := D.GetElementByID("saveButton").(*dom.HTMLButtonElement)
	saveButton.AddEventListener("click", false, func(event dom.Event) {
		Save()
	})

	clearAllButton := D.GetElementByID("clearAllButton").(*dom.HTMLButtonElement)
	clearAllButton.AddEventListener("click", false, func(event dom.Event) {
		ClearAll()
	})

	DisplayStorageContents()
}

func main() {

	switch readyState := D.ReadyState(); readyState {
	case "loading":
		D.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		println("Unexpected document.ReadyState value!")
	}
}

func Save() {

	itemKey := D.GetElementByID("itemKey").(*dom.HTMLInputElement)
	itemValue := D.GetElementByID("itemValue").(*dom.HTMLInputElement)

	if itemKey.Value == "" {
		return
	}

	SetKeyValuePair(itemKey.Value, itemValue.Value)
	itemKey.Value = ""
	itemValue.Value = ""
	DisplayStorageContents()
}

func ClearAll() {
	localStorage.Call("clear")
	DisplayStorageContents()
}

func SetKeyValuePair(itemKey string, itemValue string) {
	localStorage.Call("setItem", itemKey, itemValue)
}

func DisplayStorageContents() {

	itemList := D.GetElementByID("itemList")
	itemList.SetInnerHTML("")

	for i := 0; i < localStorage.Length(); i++ {

		itemKey := localStorage.Call("key", i)
		itemValue := localStorage.Call("getItem", itemKey)

		dtElement := D.CreateElement("dt")
		dtElement.SetInnerHTML(itemKey.String())

		ddElement := D.CreateElement("dd")
		ddElement.SetInnerHTML(itemValue.String())

		itemList.AppendChild(dtElement)
		itemList.AppendChild(ddElement)
	}

}
