package abstractions

type IFlag interface {
	GetName() string
	GetBindProperty() string

	GetValue() any
	IsVisited() bool

	SetIsVisited(bool)

	SetValueGetter(func() any)

	GetUsage() string
}
