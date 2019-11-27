package v1

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Logger(fn func(w http.ResponseWriter, r *http.Request, params httprouter.Params)) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s ", r.Method, r.URL.Path)
		fmt.Printf("%s %s ", r.Method, r.URL.Path)
		fn(w, r, params)
		log.Printf("Done in %v \n", time.Since(start))
		fmt.Printf("Done in %v \n", time.Since(start))
	}
}
