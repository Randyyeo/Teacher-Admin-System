package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/controller"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/register", teacherController.RegisterStudents).Methods("POST")
	myRouter.HandleFunc("/api/commonstudents", teacherController.GetCommonStudents).Methods("GET")
	myRouter.HandleFunc("/api/suspend", teacherController.SuspendStudent).Methods("POST")
	myRouter.HandleFunc("/api/retrievefornotification", teacherController.GetStudentsReceiveNotifications).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Successfully connected to MySQL database")

	handleRequests()
}