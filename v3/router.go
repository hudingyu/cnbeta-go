package v3

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(routes []Route) http.Handler {
	router := gin.Default()

	for _, route := range routes {
		router.GET(route.Path, route.HandleFunc)
	}
	return router
}
