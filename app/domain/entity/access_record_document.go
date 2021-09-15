package entity

import (
	"time"
)

type AccessRecordDocument struct {
	Id string `bson:"id"`
	Name  string `bson:"name"`
	Rol string `bson:"rol"`
	CurrentPlace string `bson:"current_place"`
	LastAccessTime time.Time `bson:"last_access_time"`
	Places []AccessPlace `bson:"places"`
}

func NewAccessRecordDocument(record *AccessRecord) *AccessRecordDocument{
	return &AccessRecordDocument{
		Id:             record.Id,
		Name:           record.Name,
		Rol:            record.Rol,
		CurrentPlace:   record.Place,
		LastAccessTime: record.AccessTime,
		Places:         []AccessPlace{
			NewAccessPlace(record),
		},
	}
}
