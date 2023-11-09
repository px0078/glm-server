package router

import (
	server "web-server/service"

	mux "github.com/gorilla/mux"
)

func SetRouter(r *mux.Router) {
	// user,start
	r.HandleFunc("/user/{id}", server.GetUser).Methods("GET")
	r.HandleFunc("/user", server.SignUp).Methods("POST")
	r.HandleFunc("/user", server.SignIn).Methods("GET")
	// user,end

}
