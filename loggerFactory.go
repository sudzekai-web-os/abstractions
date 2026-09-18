package abstractions

import (
	"io"

	"github.com/sudzekai-web-os/types"
)

type ILoggerFactory interface {
	SetMinLevel(types.LogLevel)
	GetMinLevel() types.LogLevel

	AddWriter() io.Writer
	GetWriters() []io.Writer

	SetPreCategory(string) ILoggerFactory
	GetPreCategory() string

	SetSubCategory(string) ILoggerFactory
	GetSubCategory() string

	Copy() ILoggerFactory

	NewLogger(string) ILogger
}
