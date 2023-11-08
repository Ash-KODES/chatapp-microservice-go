package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"chat-app-microservice/calling-service/handler"
)

func SetupRoutes(r *mux.Router, db *sql.DB) *mux.Router {
	// create,join,leave call
	r.HandleFunc("/create-call", createCallHandler(db)).Methods("POST")
	r.HandleFunc("/join-call/{callID}", joinCallHandler(db)).Methods("POST")
	r.HandleFunc("/leave-call/{callID}", leaveCallHandler(db)).Methods("POST")

	return r
}
