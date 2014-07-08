package runner

import (
	"fmt"

	"github.com/mindfork/mindfork/agent"
	"github.com/mindfork/mindfork/mindfork"
)

type Runner struct {
	agents map[string]agent.MessagingAgent
}

func Setup() (*Runner, error) {
	r, err := initRunner()
	if err != nil {
		return nil, err
	}
	// TODO: try new names until we get one that isn't taken
	mf := mindfork.New("mindfork")
	r.registerAgent(mf.Name(), mf)
	return r, nil
}

func (r *Runner) Run() {
	for _, agent := range r.agents {
		go agent.Run()
	}
	// TODO: receive messages
	// for {
	//         select {
	//         case msg := <-msgCh:
	//                 println(msg)
	//         case err := <-errCh:
	//                 fmt.Println("Error received: %s", err.Error())
	//         }
	// }
}

func initRunner() (*Runner, error) {
	// load files or whatever
	r := &Runner{}
	r.agents = make(map[string]agent.MessagingAgent)

	return r, nil
}

func (r *Runner) registerAgent(name string, a agent.MessagingAgent) (string, error) {
	if _, ok := r.agents[name]; !ok {
		r.agents[name] = a
		return name, nil
	}

	return "", fmt.Errorf("unable to register agent %q: agent already exists", name)
}
