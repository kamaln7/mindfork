package agent

type MessagingAgent interface {
	Agent
	Kill() error
}
