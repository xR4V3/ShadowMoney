package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var err error

type User struct {
	Id       int
	Username string
	Password string
	Mail     string
}

func Connection(dsn string) (*sql.DB, error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable() {
	ping()
	_, err1 := db.Query("CREATE TABLE IF NOT EXISTS users (id int NOT NULL AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), password VARCHAR(255), mail VARCHAR(255));")
	if err1 != nil {
		log.Println(err1)
	}
}

func SQLUsers(username string, password string, mail string) {
	insert, err := db.Query("INSERT INTO users (username, password, mail) VALUES ('" + username + "', '" + password + "' , '" + mail + "');")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

func GetUser(login string, password string) bool {
	sel, err := db.Query("SELECT username, password FROM users WHERE username = '" + login + "'")
	if err != nil {
		return false
	}
	for sel.Next() {
		var user User
		err = sel.Scan(&user.Username, &user.Password)
		if user.Username == login && user.Password == password {
			return true
		} else {
			return false
		}
	}

	defer sel.Close()
	return false
}

func DB(var1 *sql.DB, var2 error) {
	db = var1
	err = var2
}

func ping() {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("Data ping failed")
	} else {
		log.Println("Data ping successful")
	}
}
