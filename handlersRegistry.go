package abstractions

import (
	"net/http"

	"github.com/sudzekai-web-os/types"
)

type IHandlersRegistry interface {
	AddHandler(
		pattern string,
		hnd func(r *http.Request) (result types.HandlerResult),
	) IHandlersRegistry

	AddProtectedHandler(
		pattern string,
		hnd func(r *http.Request) (result types.HandlerResult),
		roles []string,
	) IHandlersRegistry

	AddNoFilterHandler(
		pattern string,
		hnd func(r *http.Request) (result types.HandlerResult),
	) IHandlersRegistry

	AddProtectedNoFilterHandler(
		pattern string,
		hnd func(r *http.Request) (result types.HandlerResult),
		roles []string,
	) IHandlersRegistry

	AddMiddleware(
		pos int,
		middleware types.Middleware,
	) IHandlersRegistry

	SetJwtMiddleware(jwt types.JwtMiddleware) IHandlersRegistry

	SetResultFilter(filter types.ResultFilter) IHandlersRegistry

	GetRoutes() []string
	ClearRoutes()
}
