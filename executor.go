package abstractions

import "github.com/sudzekai-web-os/types"

type IExecutor interface {
	Execute(command string, args ...string) types.CommandResult
	ExecuteInDirectory(directoryPath string, command string, args ...string) types.CommandResult
	ExecuteWithInput(input, command string, args ...string) types.CommandResult
}
