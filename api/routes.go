package api

import (
	"net/http"
	"umikyoRest/libs"
)

type routesLists struct {
	name     string
	function http.HandlerFunc
}

var routeGroups []routesLists

func app(api string, method string, fn http.HandlerFunc) {
	routeGroups = append(routeGroups, routesLists{name: api, function: libs.MethodHandler(method, fn)})
}
func RoutesMain() *http.ServeMux {
	app("/", "GET", empty)
	app("/registerUser", "POST", registerUser)
	app("/getUsers", "GET", getUsers)

	mux := http.NewServeMux()
	for _, route := range routeGroups {
		mux.HandleFunc(route.name, route.function)

	}
	return mux
}
