package routes

import "net/http"

type routesLists struct {
	name     string
	function http.HandlerFunc
}

var routeGroups []routesLists

func RoutesMain() *http.ServeMux {
	routeGroups = append(routeGroups, routesLists{name: "/", function: empty})
	routeGroups = append(routeGroups, routesLists{name: "/registerUser", function: registerUser})
	routeGroups = append(routeGroups, routesLists{name: "/getUsers", function: getUsers})

	mux := http.NewServeMux()
	for _, route := range routeGroups {
		mux.HandleFunc(route.name, route.function)

	}
	return mux
}
