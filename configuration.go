package abstractions

type IConfiguration interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool

	GetSettings() map[string]any

	SetGetString(fun func(key string) string)
	SetGetBool(fun func(key string) bool)
	SetGetInt(fun func(key string) int)
}
