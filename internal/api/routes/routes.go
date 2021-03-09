package routes

import (
	"github.com/realpamisa/RestAPI/internal/api/handler"
	"github.com/realpamisa/RestAPI/internal/api/server"
	// "github.com/gorilla/mux"
)

// type Route struct {
// 	Uri     string
// 	Method  string
// 	Handler func(w http.ResponseWriter, r *http.Request)
// }

func Load() []RouteSVC {
	routes := urlRoutes
	return routes
}

// func SetupRoutes(r *server.RestSvc) *server.RestSvc {
// 	for _, route := range Load() {
// 		r.Rest.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
// 	}
// 	return r
// }

type Route struct {
	server  server.RestSvc
	handler handler.Handler
}

const rootRoute = "/"

func Init() *Route {
	return &Route{}
}

func (r Route) SetServer(s *server.RestSvc) *Route {
	if s != nil {
		r.server = *s
	}
	return &r
}

func (r Route) SetHandler(h *handler.Handler) *Route {
	if h != nil {
		r.handler = *h
	}
	return &r
}

func (r Route) Build() *Route {
	for _, route := range Load() {
		r.server.Rest.MethodFunc(route.Method, route.Uri, route.Handler)

	}
	return &r
}
