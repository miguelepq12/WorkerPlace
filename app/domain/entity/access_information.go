package entity

import "time"

type AccessInformation struct {
	Id string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName string `json:"last_name"`
	Rol string `json:"rol"`
	SensorId string `json:"sensor_id"`
	SensorPlace string `json:"sensor_place"`
	AccessTime time.Time `json:"access_time"`
}

func NewAccessInformation() *AccessInformation {
	return &AccessInformation{}
}
