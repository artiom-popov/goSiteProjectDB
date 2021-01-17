package server

import (
	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	
	

// Db ... - глобальная переменная с подключением к БД
var Db *sqlx.DB

//InitDB функция, инициирующая подключение к БД
func InitDB() (err error) {
	//строка, содержащая данные для подключения к БД в следующем формате:
	//login:password@tcp(host:port)/dbname    3qjICBvrud 5eCmuTatIa
	var dataSourceName = "3qjICBvrud:5eCmuTatIa@tcp(remotemysql.com:3306)/3qjICBvrud"
	//подключаемся к БД, используя нужный драйвер и данные для подключения
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return
	}

	err = Db.Ping()
	return
}
