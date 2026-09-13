package abstractions

import "github.com/sudzekai/web-os-api/packages/types"

type ILogger interface {
	Log(level types.LogLevel, format string, args ...any)
	LogDebug(format string, args ...any)
	LogInformation(format string, args ...any)
	LogWarning(format string, args ...any)
	LogError(format string, args ...any)
	LogCritical(format string, args ...any)
}
