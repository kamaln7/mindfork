package main

import (
	"fmt"

	"github.com/mindfork/mindfork/mindfork"
)

func main() {
	mf := mindfork.New()
	msgCh, errCh := mf.Run()

	for {
		select {
		case msg := <-msgCh:
			println(msg)
		case err := <-errCh:
			fmt.Println("Error received: %s", err.Error())
		}
	}
}
