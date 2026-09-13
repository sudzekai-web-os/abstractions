package abstractions

type IServer interface {
	SetHost(host string)
	SetPort(port int)

	Start()
	Stop()
	WaitForShutdown()

	IsListening() bool

	GetRegistry() IHandlersRegistry
}
