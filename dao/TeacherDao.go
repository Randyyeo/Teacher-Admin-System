package TeacherDao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"myapp/db"
	"myapp/models"
	"fmt"
	"strings"
)

func RegisterStudents(request structs.RegisterStudents) (bool, error) {
	db := Db.Connect()
	students := request.Students

	if len(students) > 0 {
		var sqlStatement = "INSERT INTO teacher_registered_students VALUES "
		for _, student := range students {
			checkIfStudentIsRegisteredAndInsert(student, *db)
			sqlStatement += fmt.Sprintf("('%s', '%s'),", request.Teacher, student)
		}
		b := []byte(sqlStatement)
		b[len(b)-1] = ';'
		sqlStatement = string(b) 
		insert, err := db.Query(sqlStatement)
		
		if err != nil {
			db.Close()
			return false, err
		}
		defer insert.Close()
		db.Close()
		return true, nil
	}
	db.Close()
	return true, nil
}

func GetCommonStudents(teachers []string) ([]string, error) {
	db := Db.Connect()
	
	query := "SELECT student_email FROM teacher_registered_students WHERE teacher_email IN ("
	placeholders := strings.Trim(strings.Repeat(",?", len(teachers)), ",")
	query += placeholders + ") GROUP BY student_email HAVING COUNT(DISTINCT teacher_email) = ?"
	
	// Append the length of the teacher slice as a parameter to the query
	args := make([]interface{}, len(teachers)+1)
	for i, teacher := range teachers {
			args[i] = teacher
	}
	args[len(teachers)] = len(teachers)
	
	// Execute the query and retrieve the result set

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []string
	for rows.Next() {
			var student string
			if err := rows.Scan(&student); err != nil {
					return nil, err
			}
			students = append(students, student)
	}
	if err := rows.Err(); err != nil {
			return nil, err
	}
    
	return students, nil
}

func SuspendStudent(student string) (bool, error) {
	db := Db.Connect()
	// update student set suspend_status=true where email='student1@gmail.com';

	var sqlStatement = "UPDATE student SET suspend_status = true WHERE email='" + student + "';"
	fmt.Println(sqlStatement)
	update, err := db.Query(sqlStatement)
	defer update.Close()
	if err != nil {
		db.Close()
		return false, err
	}
	db.Close()
	return true, err
}

func GetStudentsUnderTeacher(teacher string) ([]string, error) {
	db := Db.Connect()

	query := "SELECT email from student where email in (SELECT student_email FROM teacher_registered_students WHERE teacher_email ='" + teacher + "') AND suspend_status=false;"
	rows, err := db.Query(query)
	defer rows.Close()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	var students []string
	for rows.Next() {
			var student string
			if err := rows.Scan(&student); err != nil {
					return nil, err
			}
			students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		db.Close()
		return nil, err
	}
  db.Close()
	return students, nil
}

func CheckStudentsSuspension(students []string) ([]string, error) {
	db := Db.Connect()
	query := "SELECT email from student where email in ("
	placeholders := strings.Trim(strings.Repeat(",?", len(students)), ",")
	query += placeholders + ") AND suspend_status=false;"
	args := make([]interface{}, len(students))
	for i, students := range students {
			args[i] = students
	}
	fmt.Println(args)
	
	rows, err := db.Query(query, args...)
	fmt.Println(err)
	defer rows.Close()
	if err != nil {
		db.Close()
		return nil, err
	}

	var res []string
	for rows.Next() {
			var student string
			if err := rows.Scan(&student); err != nil {
					return nil, err
			}
			res = append(res, student)
	}
	if err := rows.Err(); err != nil {
			db.Close()
			return nil, err
	}
	db.Close()
    
	return res, nil
}

func checkIfStudentIsRegisteredAndInsert(student string, db sql.DB) {
	sqlStatement := fmt.Sprintf("SELECT * from student where email='%s';", student)
	selec, err := db.Query(sqlStatement)
	defer selec.Close()
	if !selec.Next() {
		sqlStatement := fmt.Sprintf("INSERT INTO student VALUES ('%s', false)", student)
		insert, err := db.Query(sqlStatement)
		defer insert.Close()
		if err != nil {
			panic(err.Error())
		}
	
	} else if err != nil {
		panic(err.Error())
	}
	db.Close()
}

func CheckIfTeacherIsRegisteredAndInsert(teacher string) {
	db := Db.Connect()
	sqlStatement := fmt.Sprintf("SELECT * from teacher where email='%s';", teacher)
	selec, err := db.Query(sqlStatement)
	defer selec.Close()
	if !selec.Next() {
		sqlStatement := fmt.Sprintf("INSERT INTO teacher VALUES ('%s')", teacher)
		fmt.Println(sqlStatement)
		insert, err := db.Query(sqlStatement)
		defer insert.Close()
		if err != nil {
			db.Close()
		}
	
	} else if err != nil {
		panic(err.Error())
	}
	db.Close()
}