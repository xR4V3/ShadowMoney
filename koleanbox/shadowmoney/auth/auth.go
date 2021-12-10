package auth

import (
	"html/template"
	"log"
	"net/http"
)

type data struct {
	login    string
	password string
}

type status struct {
	Status string
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r) // Вывод 404
		return
	}
	home, err := template.ParseFiles("koleanbox/shadowmoney/templates/register.html")
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

func RegisterProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}

	userInfo := data{r.FormValue("login"), r.FormValue("password")}

	t, err := template.ParseFiles("koleanbox/shadowmoney/templates/index.html")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404!"))
		return
	}

	if userInfo.login == "" || userInfo.password == "" {
		istatus := status{"Заполните все поля!"}
		t.Execute(w, istatus)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		log.Printf("Лог: %s", istatus)
		return
	}
	log.Printf("Авторизация: %s %s", userInfo.login, userInfo.password)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func AuthProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}

	userInfo := data{r.FormValue("login"), r.FormValue("password")}

	t, err := template.ParseFiles("koleanbox/shadowmoney/templates/index.html")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404!"))
		return
	}

	if userInfo.login == "" || userInfo.password == "" {
		istatus := status{"Заполните все поля!"}
		t.Execute(w, istatus)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		log.Printf("Лог: %s", istatus)
		return
	}
	log.Printf("Авторизация: %s %s", userInfo.login, userInfo.password)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
