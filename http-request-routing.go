package main


import (
	"github.com/gorilla/mux"
	"net/http"
)

//   Credit: http://thenewstack.io/make-a-restful-json-api-go

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes *Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range *routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}

	return router
}