package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/spf13/viper"
)

type MySql struct {
	Conn *gorm.DB
}

func connect() *gorm.DB {
	dialect := "mysql"
	username := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "articles"

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(dialect, conn)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	return db
}

func NewMySql() *MySql {
	return &MySql{
		Conn: nil,
	}
}

func (m *MySql) Connection() *gorm.DB {
	return m.Conn
}

func (m *MySql) Connect() {
	m.Conn = connect()
}

func (m *MySql) Disconnect() {
	m.Conn.Close()
}
