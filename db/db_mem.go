package db

import "plasticine/models"

type db struct {
	groupSeq int
	groups   []models.Group

	subjectSeq int
	subjects   []models.Subject
}

func NewDB() DB {
	return &db{
		groups:   []models.Group{},
		subjects: []models.Subject{},
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

func (db *db) UpdateGroup(group models.Group) error {
	for i, g := range db.groups {
		if g.Id == group.Id {
			db.groups[i] = group
		}
	}
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

func (db db) UpdateSubject(subject models.Subject) error {
	for i, s := range db.subjects {
		if s.Id == subject.Id {
			db.subjects[i] = subject
		}
	}
	return nil
}
