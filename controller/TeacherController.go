package teacherController

import (
    "encoding/json"
    "net/http"
		"myapp/service"
		"myapp/models"
		"fmt"
)

func RegisterStudents(w http.ResponseWriter, r *http.Request) {
	var requestBody structs.RegisterStudents

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := TeacherService.RegisterStudents(requestBody)
	if result {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func GetCommonStudents(w http.ResponseWriter, r *http.Request) {

	teachers := r.URL.Query()["teacher"]

	fmt.Println(teachers)

	result, err := TeacherService.GetCommonStudents(teachers)
	fmt.Println(err)
	if err == nil {
		var students structs.CommonStudents
		students.Students = result
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(students)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func SuspendStudent(w http.ResponseWriter, r *http.Request) {
	var requestBody structs.Student

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := TeacherService.SuspendStudent(requestBody.Student)
	if result {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func GetStudentsReceiveNotifications(w http.ResponseWriter, r *http.Request) {
	var requestBody structs.TeacherNotification

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := TeacherService.GetStudentsReceiveNotifications(requestBody)
	if err == nil {
		var response structs.TeacherNotificationResponse
		response.Recipients = result
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}