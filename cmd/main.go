package main

import (
	"fmt"

	. "github.com/laof/request"
)

func main() {

	// ch := make(chan string)
	// go Direct(ch)
	// fmt.Println(<-ch)

	// cp := make(chan string)
	// go Proxy(cp)
	// fmt.Println(<-cp)

	str := Request()
	fmt.Println(str)

}
