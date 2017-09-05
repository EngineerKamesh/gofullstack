package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/EngineerKamesh/gofullstack/volume3/section1/intermediate/model"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func handleCarsData(w http.ResponseWriter, r *http.Request) {

	var cars []model.Car
	var carsDataBuffer bytes.Buffer

	dec := gob.NewDecoder(&carsDataBuffer)
	body, err := ioutil.ReadAll(r.Body)
	carsDataBuffer = *bytes.NewBuffer(body)
	err = dec.Decode(&cars)

	w.Header().Set("Content-Type", "text/plain")

	if err != nil {
		log.Println(err)
		w.Write([]byte("Something went wrong, look into it"))

	} else {
		fmt.Printf("Cars Data: %#v\n", cars)
		w.Write([]byte("Thanks, I got the slice of cars you sent me!"))
	}

}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/cars-data", handleCarsData)
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
