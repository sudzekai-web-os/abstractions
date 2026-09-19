package abstractions

type IConfigurationBuilder interface {
	AddConfigurationFile(fileName string) IConfigurationBuilder
	AddFlag(flag IFlag) IConfigurationBuilder

	Parse() IConfigurationBuilder
	LoadConfigurationFile() IConfigurationBuilder

	GetConfiguration() IConfiguration
}
