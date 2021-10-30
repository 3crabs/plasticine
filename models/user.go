package models

type UserRole string

const (
	Student = UserRole("STUDENT")
	Teacher = UserRole("TEACHER")
)

type User struct {
	Id         int      `json:"id"`
	LastName   string   `json:"last_name"`   // фамилия
	FirstName  string   `json:"first_name"`  // имя
	SecondName string   `json:"second_name"` // отчество
	Role       UserRole `json:"role"`
	GroupId    int      `json:"group_id"`
}

type UserInfo struct {
	*User
	Group `json:"group"`
}
