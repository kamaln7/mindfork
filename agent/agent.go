package agent

type Agent interface {
	Run()
	Name() string
}
