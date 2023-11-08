package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"chat-app-microservice/calling-service/handler"
)

func SetupRoutes(r *mux.Router, db *sql.DB) *mux.Router {
	// create,join,leave call
	r.HandleFunc("/create-call", handler.CreateCallHandler(db)).Methods("POST")
	r.HandleFunc("/join-call/{callID}", handler.JoinCallHandler(db)).Methods("POST")
	r.HandleFunc("/leave-call/{callID}", handler.LeaveCallHandler(db)).Methods("POST")

	return r
}
