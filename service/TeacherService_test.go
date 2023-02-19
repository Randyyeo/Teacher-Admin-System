package TeacherService

import (
	"testing"
	"myapp/models"
)

type MockTeacherService struct {
	RegisterStudents     func(request structs.RegisterStudents) (bool, error)
	GetCommonStudents func(teachers []string) ([]string, error)
}

func (m *MockTeacherService) RegisterStudents(request structs.RegisterStudents) (bool, error) {
	return m.RegisterStudents(request)
}

func (m *MockTeacherService) GetCommonStudents(teachers [string]) ([]string, error) {
	return m.GetCommonStudents(teachers)
}

func Test

