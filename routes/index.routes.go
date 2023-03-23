package routes

import "net/http"

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Home Handler"))
}
