package main

import (
	"ShadowMoney/koleanbox/shadowmoney/auth"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // Вывод 404
		return
	}

	home, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
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
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
