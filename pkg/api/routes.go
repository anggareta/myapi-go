package api

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"baca semua",
		"GET",
		"/baca",
		baca,
	},
	Route{
		"baca satu",
		"GET",
		"/bacasatu/{id}",
		bacasatu,
	},
	Route{
		"test-post",
		"POST",
		"/objek",
		objek,
	},
	Route{
		"test-SP",
		"POST",
		"/tulis",
		tulis,
	},
}
