package request

import "fmt"

func Request() string {

	nd := make(chan string)
	np := make(chan string)

	go Direct(nd)
	go Proxy(np)

	direct := <-nd
	proxy := <-np

	nodes := ""
	if direct != "" {
		nodes = direct
	} else if proxy != "" {
		nodes = proxy
	} else {
		fmt.Println("Sorry, connection failed please try again")
	}
	return nodes
}
