package mindfork

import (
	"fmt"
	"os"
	"time"
)

type Mindfork struct {
	errCh     chan error
	msgCh     chan string
	commandCh chan Command

	name string

	home  string
	files []string
}

func New(name string) *Mindfork {
	mf := &Mindfork{
		home:      os.Getenv("HOME") + "/.mindfork",
		files:     []string{},
		name:      name,
		errCh:     make(chan error),
		msgCh:     make(chan string),
		commandCh: make(chan Command),
	}

	return mf
}

func (mf *Mindfork) Run() {
	go mf.run()
}

func (mf *Mindfork) Name() string {
	return mf.name
}

func (mf *Mindfork) Kill() error {
	if _, err := mf.sendCommand(Command{kill}); err != nil {
		return fmt.Errorf("kill failed: %s", err.Error())
	}

	return nil
}

func (mf *Mindfork) sendCommand(command Command) (bool, error) {
	mf.commandCh <- command
	// TODO: receive result (pass/fail?) and error
	return true, nil
}

func (mf *Mindfork) run() {
	// TODO: plan something meaningful using statement of intent + children
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
			mf.errCh <- fmt.Errorf("%d", i)
		default:
			mf.msgCh <- fmt.Sprintf("current %d\nNow: %s", i, time.Now().String())
		}
	}
}
