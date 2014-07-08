package main

import "github.com/mindfork/mindfork/runner"

func main() {
	r, err := runner.Setup()
	if err != nil {
		panic(err)
	}

	r.Run()
}
