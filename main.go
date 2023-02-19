package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/handlers"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/register", teacherHandler.RegisterStudents).Methods("POST")
	myRouter.HandleFunc("/api/commonstudents", teacherHandler.GetCommonStudents).Methods("GET")
	myRouter.HandleFunc("/api/suspend", teacherHandler.SuspendStudent).Methods("POST")
	myRouter.HandleFunc("/api/retrievefornotification", teacherHandler.GetStudentsReceiveNotifications).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Successfully connected to MySQL database")

	handleRequests()
}