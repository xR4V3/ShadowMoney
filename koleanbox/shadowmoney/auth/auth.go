package auth

import (
	"log"
	"net/http"
)

func Auth() {

}

func AuthProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}

	log.Println("Логгинг...")
}
