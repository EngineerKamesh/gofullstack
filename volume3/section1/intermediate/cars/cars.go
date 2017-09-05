package main

import (
	"bytes"
	"encoding/gob"
	"html/template"

	"github.com/EngineerKamesh/gofullstack/volume3/section1/intermediate/model"
	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

const CarItemTemplate = `
	<td class="mdl-data-table__cell--non-numeric">{{.ModelName}}</td>
	<td class="mdl-data-table__cell--non-numeric">{{.Color}}</td>
	<td class="mdl-data-table__cell--non-numeric">{{.Manufacturer}}</td>
`

var D = dom.GetWindow().Document()

func main() {

	nano := model.Car{ModelName: "Nano", Color: "Yellow", Manufacturer: "Tata"}
	ambassador := model.Car{ModelName: "Ambassador", Color: "White", Manufacturer: "HM"}
	omni := model.Car{ModelName: "Omni", Color: "Red", Manufacturer: "Maruti Suzuki"}
	cars := []model.Car{nano, ambassador, omni}

	println("Cars Template Example")

	autoTableBody := D.GetElementByID("autoTableBody")
	for i := 0; i < len(cars); i++ {
		trElement := D.CreateElement("tr")
		tpl := template.New("template")
		tpl.Parse(CarItemTemplate)
		var buff bytes.Buffer
		tpl.Execute(&buff, cars[i])
		trElement.SetInnerHTML(buff.String())
		autoTableBody.AppendChild(trElement)
	}

	var carsDataBuffer bytes.Buffer
	enc := gob.NewEncoder(&carsDataBuffer)
	enc.Encode(cars)

	xhrResponse, err := xhr.Send("POST", "/cars-data", carsDataBuffer.Bytes())

	if err != nil {
		println(err)
	}

	println("xhrResponse: ", string(xhrResponse))

}
