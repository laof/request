package request

import "fmt"

func Request() string {

	nodes := Direct()

	if nodes == "" {
		fmt.Println("connection fail, starting proxy channel...")
		nodes = Proxy()
	}

	return nodes
}
