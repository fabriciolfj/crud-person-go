package datasource

import (
	"fmt"
	"log"

	"github.com/magiconair/properties"
	"github.com/person/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		p.MustGetString("db_user"),
		p.MustGetString("db_password"),
		p.MustGetString("db_host"),
		p.MustGetString("db_port"),
		p.MustGetString("db_name"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil || DB == nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	err = DB.AutoMigrate(&model.Person{})

	if err != nil {
		log.Printf("Erro ao realizar a migração: %v", err)
	}

	log.Println("Conexão estabelecida", DB.Name())
	log.Println("DB Pointer:", DB)
}
