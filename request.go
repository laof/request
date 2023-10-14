package request

import (
	"log"
)

func Request() string {

	log.Println("connection ...")
	nodes := Direct()

	if nodes == "" {
		log.Println("starting proxy channel ...")
		nodes = Proxy()
	}

	return nodes
}
