package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/alicobanserver/controllers"
	"github.com/gorilla/mux"
)

func Routes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/set", controllers.Set).Methods("POST")
	router.HandleFunc("/get", controllers.Get).Methods("GET")
	router.HandleFunc("/fetch", controllers.Fetch).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(http.ListenAndServe(":4444", router))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, router))
	}

}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Welcome To Getir Case API"}`))

	return
}
