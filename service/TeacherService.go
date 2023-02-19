package TeacherService

import (
	"myapp/models"
	"myapp/dao"
	"fmt"
	"strings"
)

func RegisterStudents(request structs.RegisterStudents) (bool, error) {

	TeacherDao.CheckIfTeacherIsRegisteredAndInsert(request.Teacher)
	status, err := TeacherDao.RegisterStudents(request)
	if err != nil {
		return false, err
	}
	return status, nil
}

func GetCommonStudents(teachers []string) ([]string, error) {
	return TeacherDao.GetCommonStudents(teachers)
}

func SuspendStudent(student string) (bool, error) {
	return TeacherDao.SuspendStudent(student)
}

func GetStudentsReceiveNotifications(request structs.TeacherNotification) ([]string, error) {

	studentsUnderTeacher, err := TeacherDao.GetStudentsUnderTeacher(request.Teacher)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	checkStudentsSpecifiedSuspension := []string{}
	notification := request.Notification
	stringArr := strings.Split(notification, " ")
	for i:=0; i<len(stringArr); i++ {
		if strings.Contains(stringArr[i], "@") && stringNotInArray(stringArr[i][1:], studentsUnderTeacher){
			newString := stringArr[i][1:]
			checkStudentsSpecifiedSuspension = append(checkStudentsSpecifiedSuspension, newString)
		}
	}

	studentsSpecified, err := TeacherDao.CheckStudentsSuspension(checkStudentsSpecifiedSuspension)
	result := append(studentsUnderTeacher, studentsSpecified...)

	return result, err
}

func stringNotInArray(student string, studentArr []string) bool {
	for i:=0;i<len(studentArr);i++ {
		if studentArr[i] == student {
			return false
		}
	}
	return true
}