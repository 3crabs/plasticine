package db

import "plasticine/models"

type db struct {
	groups []models.Group
}

func NewDB() DB {
	return &db{groups: []models.Group{}}
}

func (db *db) AddGroups(groups []models.Group) error {
	for _, g := range groups {
		db.groups = append(db.groups, g)
	}
	return nil
}

func (db *db) GetGroups() []models.Group {
	return db.groups
}
