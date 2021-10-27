package db

import "plasticine/models"

type db struct {
	seq    int
	groups []models.Group
}

func NewDB() DB {
	return &db{groups: []models.Group{}}
}

func (db *db) AddGroup(group models.Group) error {
	db.seq++
	group.Id = db.seq
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
