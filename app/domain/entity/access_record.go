package entity

import (
	"encoding/json"
	"errors"
	"time"
)

type AccessRecord struct {
	Id string `json:"id"`
	Name  string `json:"name"`
	Rol string `json:"rol"`
	Place string `json:"place"`
	AccessTime time.Time `json:"access_time"`
}

func (r *AccessRecord) DecodeByte(value []byte) error {
	err :=json.Unmarshal(value,r)
	if err != nil {
		return errors.New("input is not an access record entity")
	}
	return nil
}
