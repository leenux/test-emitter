package main

import (
	"log"
	"time"

	emitter "github.com/emitter-io/go/v2"
)

//channel v1/#/

func main() {
	clientA()
	clientB()

	// stop after 10 seconds
	time.Sleep(1 * time.Second)
}

func clientA() {
	const key = "GTYN24PdGhiA1yPODeimI32o49yOiayX" // read on $public/down/

	c, _ := emitter.Connect("tcp://127.0.0.1:8080", func(_ *emitter.Client, msg emitter.Message) {
		log.Printf("[emitter] -> [A] received: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	})

	log.Println("[emitter] -> [A] my name is " + c.ID())

	log.Println("[emitter] <- [A] subscribing to 'v1/#/'")
	if err := c.Subscribe(key, "v1/", nil); err != nil {
		log.Printf("[emitter] -> [A] Subscribe error:%s", err.Error())
	}
}

func clientB() {
	const key = "GTYN24PdGhiA1yPODeimI32o49yOiayX" // everything on $public/down/

	c, _ := emitter.Connect("tcp://127.0.0.1:8080", func(_ *emitter.Client, msg emitter.Message) {
		log.Printf("[emitter] -> [B] received: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	})

	log.Println("[emitter] -> [B] my name is " + c.ID())

	log.Println("[emitter] <- [B] subscribing to 'v1/1122/attr/'")
	if err := c.Subscribe(key, "v1/1122/attr/", func(_ *emitter.Client, msg emitter.Message) {
		log.Printf("[emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	}); err != nil {
		log.Printf("[emitter] -> [B] Subscribe error:%s", err.Error())
	}

	log.Println("[emitter] <- [B] publishing to 'v1/1122/attr/'")
	if err := c.Publish(key, "v1/1122/attr/", "{hello: \"world\"}"); err != nil {
		log.Printf("[emitter] -> [B] Publish error:%s", err.Error())
	}
}
