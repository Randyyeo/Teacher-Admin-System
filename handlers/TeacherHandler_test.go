package teacherHandler

import (
	"testing"
	"myapp/models"
)

type MockTeacherService struct {
	RegisterStudentsFunc     func(request structs.RegisterStudents) (bool, error)
	GetCommonStudentsFunc func(teachers []string) ([]string, error)
}

func (m *MockTeacherService) RegisterStudents(request structs.RegisterStudents) (bool, error) {
	return m.RegisterStudents(request)
}

func (m *MockTeacherService) GetCommonStudents(teachers []string) ([]string, error) {
	return m.GetCommonStudents(teachers)
}

func TestRegisterStudentsHandler(t *testing.T) {
	
}

