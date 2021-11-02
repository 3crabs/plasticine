package db

import "plasticine/models"

type DB interface {
	AddGroup(group models.Group) error
	GetGroups() []models.Group
	GetGroup(groupId int) (*models.Group, error)
	UpdateGroup(group models.Group) error
	DeleteGroup(groupId int) error
	GetGroupStudents(groupId int) []models.User

	AddSubject(subject models.Subject) error
	GetSubjects() []models.Subject
	UpdateSubject(subject models.Subject) error
	GetSubject(subjectId int) (*models.Subject, error)

	AddUser(user models.User) error
	UpdateUser(user models.User) error
	GetUsersByRole(role models.UserRole) []models.User
	GetUserInfo(studentId int) (*models.UserInfo, error)
	DeleteUser(userId int) error
}
