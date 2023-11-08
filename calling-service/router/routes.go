// router/routes.go
// Define the routes for video calling

package router

import (
    "github.com/gorilla/mux"
    "net/http"
    "database/sql"
)

func SetupVideoCallingRoutes(r *mux.Router, db *sql.DB) *mux.Router {
    // Define video calling routes
    r.HandleFunc("/create-call", createCallHandler(db)).Methods("POST")
    r.HandleFunc("/join-call/{callID}", joinCallHandler(db)).Methods("POST")
    r.HandleFunc("/leave-call/{callID}", leaveCallHandler(db)).Methods("POST")

    return r
}
