package service

import (
	"github.com/person/datasource"
	"github.com/person/model"

	"fmt"
)

func Save(p *model.Person) error {
	if datasource.DB == nil {
		return fmt.Errorf("DB is not initialized")
	}

	person := p.GenerateUUID()
	result := datasource.DB.Create(person)
	return result.Error
}
