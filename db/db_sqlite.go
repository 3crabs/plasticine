package db

import (
	"plasticine/models"
)

type dbSQLite struct{}

func newDbSQLite() DB {
	return &dbSQLite{}
}

func (db *dbSQLite) AddGroup(group models.Group) error {
	panic("implement me")
}

func (db *dbSQLite) GetGroups() []models.Group {
	panic("implement me")
}

func (db *dbSQLite) GetGroup(groupId int) (*models.GroupInfo, error) {
	panic("implement me")
}

func (db *dbSQLite) UpdateGroup(group models.Group) error {
	panic("implement me")
}

func (db *dbSQLite) DeleteGroup(groupId int) error {
	panic("implement me")
}

func (db *dbSQLite) GetGroupStudents(groupId int) []models.User {
	panic("implement me")
}

func (db *dbSQLite) AddSubject(subject models.Subject) error {
	panic("implement me")
}

func (db *dbSQLite) GetSubjects() []models.Subject {
	panic("implement me")
}

func (db *dbSQLite) UpdateSubject(subject models.Subject) error {
	panic("implement me")
}

func (db *dbSQLite) GetSubject(subjectId int) (*models.Subject, error) {
	panic("implement me")
}

func (db *dbSQLite) DeleteSubject(subjectId int) error {
	panic("implement me")
}

func (db *dbSQLite) AddUser(user models.User) error {
	panic("implement me")
}

func (db *dbSQLite) UpdateUser(user models.User) error {
	panic("implement me")
}

func (db *dbSQLite) GetUsersByRole(role models.UserRole) []models.User {
	panic("implement me")
}

func (db *dbSQLite) GetUserInfo(studentId int) (*models.UserInfo, error) {
	panic("implement me")
}

func (db *dbSQLite) DeleteUser(userId int) error {
	panic("implement me")
}

func (db *dbSQLite) AddPlace(place models.Place) error {
	panic("implement me")
}

func (db *dbSQLite) GetPlaces() []models.Place {
	panic("implement me")
}

func (db *dbSQLite) UpdatePlace(place models.Place) error {
	panic("implement me")
}

func (db *dbSQLite) DeletePlace(placeId int) error {
	panic("implement me")
}

func (db *dbSQLite) AddLesson(lesson models.Lesson) error {
	panic("implement me")
}

func (db *dbSQLite) GetLessons(groupId int) []models.Lesson {
	panic("implement me")
}

func (db *dbSQLite) UpdateLesson(lesson models.Lesson) error {
	panic("implement me")
}

func (db *dbSQLite) DeleteLesson(lessonId int) error {
	panic("implement me")
}
