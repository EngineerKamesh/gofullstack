package main

import (
	"flag"
	"log"
	"net/url"
	"time"

	"github.com/gopherjs/websocket"
	"honnef.co/go/js/dom"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan bool, 1)

	dom.GetWindow().AddEventListener("beforeunload", false, func(event dom.Event) {
		interrupt <- true
	})

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, err := websocket.Dial(u.String())
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			message := make([]byte, 1024)
			var n int
			n, err = c.Read(message)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message[:n])
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			_, err = c.Write([]byte(t.String()))

			if err != nil {
				log.Println("write:", err)
				return
			}

		case <-interrupt:
			c.Close()
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return

		}
	}
}
