package auth

import (
	"golangify.com/snippetbox/koleanbox/shadowmoney/database"
	"html/template"
	"log"
	"net/http"
)

type dataAuth struct {
	login    string
	password string
}

type dataReg struct {
	login     string
	password1 string
	password2 string
	mail      string
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

	userInfo := dataReg{r.FormValue("username"), r.FormValue("password1"), r.FormValue("password2"), r.FormValue("mail")}

	t, err := template.ParseFiles("koleanbox/shadowmoney/templates/register.html")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404!"))
		return
	}

	if userInfo.login == "" || userInfo.password1 == "" || userInfo.password2 == "" || userInfo.mail == "" {
		istatus := status{"Заполните все поля!"}
		t.Execute(w, istatus)
		http.Redirect(w, r, "/register", http.StatusMovedPermanently)
		log.Printf("Лог: %s", istatus)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	database.SQLUsers(userInfo.login, userInfo.password1, userInfo.mail)
	log.Printf("Регистрация: %s %s", userInfo.login, userInfo.mail)
}

func AuthProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}

	userInfo := dataAuth{r.FormValue("login"), r.FormValue("password")}

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
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	if !database.GetUser(userInfo.login, userInfo.password) {
		log.Printf("Неверные данные: %s", userInfo.login)
	} else {
		log.Printf("Авторизация: %s", userInfo.login)
	}
}
