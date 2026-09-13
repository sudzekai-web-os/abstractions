package abstractions

import (
	"io"

	"github.com/sudzekai-web-os/types"
)

type ILoggerFactory interface {
	SetWriter(writer io.Writer)
	SetMinLevel(level types.LogLevel)
	SetMinLevelStr(level string) error

	GetWriter() io.Writer
	GetMinLevel() types.LogLevel

	NewLogger(category string) ILogger
}
