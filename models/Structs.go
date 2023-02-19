package structs

type Student struct {
	Student string `json:"student"`
}

type TeacherNotification struct {
	Teacher string `json:"teacher"`
	Notification string `json:"notification"`
}

type RegisterStudents struct {
	Teacher  string  `json:"teacher"`
	Students []string `json:"students"`
}

type CommonStudents struct {
	Students []string `json:"students"`
}

type TeacherNotificationResponse struct {
	Recipients []string `json:"recipients"`
}