package main

import "os"

type mindfork struct {
	home  string
	files []string
}

func main() {
	// if ~/.mindfork not defined, define it
	//

	mf := createNewMindfork()
	mf.run()
}

func createNewMindfork() *mindfork {
	return &mindfork{
		home:  os.Getenv("HOME") + "/.mindfork",
		files: []string{},
	}
}

func (mf *mindfork) run() {
	println("foo")
}
