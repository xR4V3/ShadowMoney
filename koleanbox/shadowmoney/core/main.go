package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golangify.com/snippetbox/koleanbox/shadowmoney/auth"
	"golangify.com/snippetbox/koleanbox/shadowmoney/database"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // Вывод 404
		return
	}
	home, err := template.ParseFiles("koleanbox/shadowmoney/templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 404)
		return
	}
	err = home.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/authProcess", auth.AuthProcess)
	mux.HandleFunc("/register", auth.Register)
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	dsn := "u641154_sm:shadowmoney@tcp(185.179.190.245:3306)/u641154_sm"
	db, err := database.Connection(dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("База данных подключена!")
	}
	defer db.Close()
	err1 := http.ListenAndServe(":8080", mux)
	log.Fatal(err1)
}
