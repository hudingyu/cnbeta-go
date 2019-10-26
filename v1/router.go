package v1

import "github.com/julienschmidt/httprouter"

func NewRouter(routes []Route) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, Logger(Cors(route.HandleFunc)))
	}
	return router
}
