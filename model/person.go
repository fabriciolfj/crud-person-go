package model

type Person struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
	Age  int
}

func (Person) TableName() string {
	return "persons"
}
