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

	result := datasource.DB.Create(p)
	return result.Error
}
