package abstractions

type IModule interface {
	Name() string
	Version() string
	Description() string
	Initialize(
		registry IHandlersRegistry,
		loggerFactory ILoggerFactory,
		executor IExecutor) error
}
