# Teacher Admin System

This application is an admin System for teachers where they can perform administrative functions for their students. Teachers and students are identified by their email addresses. The 4 main functions that teachers can perform are:
- Registering students to a specified teacher
- Retrieving a list of common students given a list of teachers
- Suspend a specified student
- Retrieve a list of students who can receive a notification


## Design Architecture

The application is designed using an MVC design pattern. The view portion comes from the `main.go` which is where the application is running from. The different functions in `TeacherController` will be called based on the endpoint called by the user, for e.g. `/api/register`. By the user's interaction with the View, different service functions will be called. For example, when the user calls an API endpoint for e.g. `/api/register`, the `TeacherService.RegisterStudents` will be called where the business logic lies within it. It then updates the database through the dao layer, `TeacherDao`.

In the project, it contains 5 folders:
1. controller -> how the endpoint will be handled
2. dao -> layer to interact with the database
3. db -> connection to the database
4. models -> different structs used in the application
5. service -> service layer to interact with the dao layer

## Setting up environment

Please ensure that you have Go installed and updated to the latest version. You can check by typing this:
```
go version
```

There are certain third-party libraries that are used in these applications. They are specified in the `go.mod` and `go.sum` files. They should automatically be installed once you run the application. If in the case this does not happen, please run the following commands:
```
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux v1.8.0
```

## Database Setup
You will need a .env file in order to run this application correctly. Please correct one and add the following variables to your .env file. It should be in the root folder of the project. Please ensure that your `DB_NAME` is the same as the database you have created on your local machine. 
```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

To run the project, use the following command in the terminal:
```
go run main.go
```

## Using
To try the different endpoints, simply use postman or whatever other tool you have, to call the endpoints.
### Register Students
```
/api/register -> POST
Request Body:
{
    "teacher": "teacher5@gmail.com",
    "students": []
}
```

### Get Common Students
```
/api/commonstudents?teacher=teacher@gmail.com -> GET
```

### Suspend
```
/api/suspend -> POST
Request Body:
{
    "student": "student1@gmail.com"
}
```

### Retrieve student list for notification
```
/api/retrievefornotification -> POST
Request Body:
{
    "teacher": "teacher5@gmail.com",
    "notification": "Hello @student1@gmail.com @student2@gmail.com"
}
```
## Tests
Unfortunately, due to the way I have designed the application, I am unable to mock the different dependencies for each file and test them. For future considerations, I would have binded the dependencies into the relevant functions so that I would be able to mock them.

## Future Considerations

There are some considerations that have not been implemented in this project
1. Returning relevant messages when the student/teacher is not in the database. For example, when a teacher wants to suspend a student, the application should return a 404 status, mentioning that the user has not been admitted in the system
2. I would change my design and bind the dependencies to each file in order to be able to test them.
3. Add integration tests to test the endpoints

