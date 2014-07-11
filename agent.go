package mindfork/agent

type Agent interface {
	Run() error
	Kill() error
	SendMessage(string) error
}
