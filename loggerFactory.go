package abstractions

import (
	"github.com/sudzekai-web-os/types"
)

type ILoggerFactory interface {
	SetMinLevel(types.LogLevel)
	GetMinLevel() types.LogLevel

	AddWriter(ILoggerWriter) ILoggerFactory
	GetWriters() []ILoggerWriter

	SetPreCategory(string) ILoggerFactory
	GetPreCategory() string

	SetSubCategory(string) ILoggerFactory
	GetSubCategory() string

	Copy() ILoggerFactory

	NewLogger(string) ILogger
}
