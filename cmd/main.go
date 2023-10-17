package main

import (
	"fmt"
	"log"

	"github.com/laof/request"
)

func main() {

	str, err := request.New()

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	fmt.Println(str)

}
