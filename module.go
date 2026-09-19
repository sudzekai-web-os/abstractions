package abstractions

type IModule interface {
	Name() string
	Version() string
	Description() string
	Start() error
}

type ILoggerFactoryConsumer interface {
	AddLoggerFactory(
		loggerFactory ILoggerFactory,
	)
}

type IExecutorConsumer interface {
	AddExecutor(
		executor IExecutor,
	)
}

type IHandlersRegistryConsumer interface {
	AddHandlersRegistry(
		registry IHandlersRegistry,
	)
}

type IConfigurationConsumer interface {
	AddConfiguration(
		configuration IConfiguration,
	)
}
