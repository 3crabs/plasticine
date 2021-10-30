package db

import "plasticine/models"

type DB interface {
	AddGroup(group models.Group) error
	GetGroups() []models.Group
	UpdateGroup(group models.Group) error

	AddSubject(subject models.Subject) error
	GetSubjects() []models.Subject
	UpdateSubject(subject models.Subject) error

	AddUser(user models.User) error
	UpdateUser(user models.User) error
	GetUsersByRole(role models.UserRole) []models.User
}
