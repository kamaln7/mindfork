package mindfork

import (
	"fmt"
	"os"
	"time"
)

type Mindfork struct {
	home  string
	files []string
}

func New() *Mindfork {
	mf := &Mindfork{
		home:  os.Getenv("HOME") + "/.mindfork",
		files: []string{},
	}

	return mf
}

func (mf *Mindfork) Run() (chan string, chan error) {
	errCh := make(chan error)
	msgCh := make(chan string)

	go mf.run(msgCh, errCh)

	return msgCh, errCh
}

func (mf *Mindfork) run(msg chan string, errch chan error) {
	for i := 0; true; i++ {
		divTen := i%10 == 0
		divHundred := i%100 == 0
		divThousand := i%1000 == 0

		switch {
		case divThousand:
			fallthrough
		case divHundred:
			fallthrough
		case divTen:
			errch <- fmt.Errorf("%d", i)
		default:
			msg <- fmt.Sprintf("current %d\nNow: %s", i, time.Now().String())
		}
	}
}
