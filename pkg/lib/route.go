package lib

import (
	"net/http"

	"go.uber.org/fx"
)

type Route interface {
	http.Handler
	// Pattern reports the path at which this is registered.
	Pattern() string
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}
