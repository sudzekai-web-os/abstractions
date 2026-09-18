package abstractions

import "github.com/sudzekai-web-os/types"

type ILoggerWriter interface {
	Write(types.LogEntry)
	WriteBatch([]types.LogEntry)
}
