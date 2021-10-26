package db

import "plasticine/models"

type DB interface {
	AddGroups(groups []models.Group) error
	GetGroups() []models.Group
}
