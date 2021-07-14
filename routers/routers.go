package routers

import "net/http"

type Router struct {
	Get    http.Handler
	Post   http.Handler
	Put    http.Handler
	Delete http.Handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if router.Get != nil {
			router.Get.ServeHTTP(w, r)
		}
	case http.MethodPost:
		if router.Post != nil {
			router.Post.ServeHTTP(w, r)
		}
	case http.MethodPut:
		if router.Put != nil {
			router.Put.ServeHTTP(w, r)
		}
	case http.MethodDelete:
		if router.Delete != nil {
			router.Delete.ServeHTTP(w, r)
		}
	}
}
