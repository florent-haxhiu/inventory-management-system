package model

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	UserId      uuid.UUID `json:"userId"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

type User struct {
	Id          uuid.UUID   `json:"id"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Username    string      `json:"username"`
	Items       []Inventory `json:"items"`
	CreatedTime time.Time   `json:"createdTime"`
	UpdatedTime time.Time   `json:"updatedTime"`
}
