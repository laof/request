package request

import "fmt"

func Request() string {

	nodes := Direct()

	if nodes == "" {
		fmt.Println("starting proxy channel")
		nodes = Proxy()
	}

	return nodes
}
