package db

import "plasticine/models"

type DB interface {
	AddGroup(group models.Group) error
	GetGroups() []models.Group
	UpdateGroup(group models.Group) error
}
