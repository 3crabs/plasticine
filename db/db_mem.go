package db

import (
	"errors"
	"plasticine/models"
)

type db struct {
	groupSeq int
	groups   []models.Group

	subjectSeq int
	subjects   []models.Subject

	userSeq int
	users   []models.User
}

func NewDB() DB {
	return &db{
		groupSeq: 0,
		groups:   []models.Group{},

		subjectSeq: 0,
		subjects:   []models.Subject{},

		userSeq: 0,
		users:   []models.User{},
	}
}

func (db *db) AddGroup(group models.Group) error {
	db.groupSeq++
	group.Id = db.groupSeq
	db.groups = append(db.groups, group)
	return nil
}

func (db *db) GetGroups() []models.Group {
	return db.groups
}

func (db *db) GetGroup(groupId int) (*models.Group, error) {
	for _, g := range db.groups {
		if g.Id == groupId {
			return &g, nil
		}
	}
	return nil, errors.New("group not found")
}

func (db *db) UpdateGroup(group models.Group) error {
	for i, g := range db.groups {
		if g.Id == group.Id {
			db.groups[i] = group
		}
	}
	return nil
}

func (db *db) DeleteGroup(groupId int) error {
	var groups []models.Group
	for _, group := range db.groups {
		if group.Id != groupId {
			groups = append(groups, group)
		}
	}
	db.groups = groups
	return nil
}

func (db *db) AddSubject(subject models.Subject) error {
	db.subjectSeq++
	subject.Id = db.subjectSeq
	db.subjects = append(db.subjects, subject)
	return nil
}

func (db *db) GetSubjects() []models.Subject {
	return db.subjects
}

func (db *db) UpdateSubject(subject models.Subject) error {
	for i, s := range db.subjects {
		if s.Id == subject.Id {
			db.subjects[i] = subject
		}
	}
	return nil
}

func (db *db) GetSubject(subjectId int) (*models.Subject, error) {
	for _, s := range db.subjects {
		if s.Id == subjectId {
			return &s, nil
		}
	}
	return nil, errors.New("group not found")
}

func (db *db) DeleteSubject(subjectId int) error {
	var subjects []models.Subject
	for _, s := range db.subjects {
		if s.Id != subjectId {
			subjects = append(subjects, s)
		}
	}
	db.subjects = subjects
	return nil
}

func (db *db) GetUsersByRole(role models.UserRole) []models.User {
	var users []models.User
	for _, u := range db.users {
		if u.Role == role {
			users = append(users, u)
		}
	}
	if users == nil {
		return []models.User{}
	}
	return users
}

func (db *db) AddUser(user models.User) error {
	db.userSeq++
	user.Id = db.userSeq
	db.users = append(db.users, user)
	return nil
}

func (db *db) UpdateUser(user models.User) error {
	for i, u := range db.users {
		if u.Id == user.Id {
			db.users[i] = user
		}
	}
	return nil
}

func (db *db) GetUserInfo(userId int) (*models.UserInfo, error) {
	var user *models.User
	for _, u := range db.users {
		if u.Id == userId {
			user = &u
		}
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	userInfo := models.UserInfo{
		User: user,
	}
	if user.GroupId != 0 {
		group, err := db.GetGroupById(user.GroupId)
		if err != nil {
			return nil, err
		}
		userInfo.Group = group
	}
	return &userInfo, nil
}

func (db *db) GetGroupById(groupId int) (*models.Group, error) {
	var group *models.Group
	for _, g := range db.groups {
		if g.Id == groupId {
			group = &g
		}
	}
	if group == nil {
		return nil, errors.New("group not found")
	}
	return group, nil
}

func (db *db) GetGroupStudents(groupId int) []models.User {
	var students []models.User
	for _, user := range db.users {
		if user.Role == models.Student && user.GroupId == groupId {
			students = append(students, user)
		}
	}
	return students
}

func (db *db) DeleteUser(userId int) error {
	var users []models.User
	for _, user := range db.users {
		if user.Id != userId {
			users = append(users, user)
		}
	}
	db.users = users
	return nil
}
