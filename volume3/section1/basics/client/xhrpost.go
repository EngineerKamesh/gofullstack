package main

import (
	"encoding/json"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

func lowercaseTextTransformer() {
	d := dom.GetWindow().Document()
	textToLowercase := d.GetElementByID("textToLowercase").(*dom.HTMLInputElement)

	textBytes, err := json.Marshal(textToLowercase.Value)
	if err != nil {
		println("Encountered error while attempting to marshal JSON: ", err)
		println(err)
	}

	data, err := xhr.Send("POST", "/lowercase-text", textBytes)
	if err != nil {
		println("Encountered error while attempting to submit POST request via XHR: ", err)
		println(err)
	}

	var s string
	err = json.Unmarshal(data, &s)

	if err != nil {
		println("Encountered error while attempting to umarshal JSON data: ", err)
	}
	textToLowercase.Set("value", s)
}
