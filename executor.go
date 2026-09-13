package abstractions

import "github.com/sudzekai/web-os-api/packages/types"

type IExecutor interface {
	Execute(command string, args ...string) types.CommandResult
	ExecuteInDirectory(directoryPath string, command string, args ...string) types.CommandResult
}
