package entity

import "time"

type AccessPlace struct {
	PlaceName string `bson:"place_name"`
	InDate time.Time `bson:"in_date"`
	OutDate time.Time `bson:"out_date"`
	DurationInSecond float64 `bson:"duration_in_second"`
}

func NewAccessPlace(record *AccessRecord) AccessPlace {
	return AccessPlace{
		PlaceName:        record.Place,
		InDate:           record.AccessTime,
		OutDate:          time.Time{},
		DurationInSecond: 0,
	}
}