package model

import (
	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
	Age  int
	Uuid string
}

func (Person) TableName() string {
	return "persons"
}

func (p Person) GenerateUUID() Person {
	p.Uuid = uuid.NewV4().String()
	return p
}
