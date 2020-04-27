package webapp

import (
	"encoding/json"
	"fmt"
	"github.com/dbut2/shortener/pkg/shortener"
	"github.com/go-chi/chi"
	"math/rand"
	"net/http"
	"time"
)

func file(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./web"))
	fs.ServeHTTP(w, r)
}
func shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.Form.Get("code")
	url := r.Form.Get("url")

	if url == "" {
		respondError(w, http.StatusBadRequest, "missing url")
		return
	}
	if code == "" {
		rand.Seed(time.Now().UnixNano())
		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
		for i := 0; i < 4; i++ {
			code += string(chars[rand.Intn(len(chars))])
		}
	}

	if shortener.CodeExists(code) {
		respondError(w, http.StatusBadRequest, "code taken")
		return
	}

	err := shortener.Shorten(url, code)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	respondJSON(w, http.StatusOK, fmt.Sprintf("{\"code\":\"%s\"}", code))

}

func lengthen(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := chi.URLParam(r, "code")

	if !shortener.CodeExists(code) {
		respondStatus(w, http.StatusNotFound)
	}

	url, err := shortener.Lengthen(code)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	respondRedirect(w, url)
}

func respondStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, fmt.Sprintf("{status:%d,message:\"%s\"}", status, message))
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	respondWith, err := json.Marshal(payload)
	if err != nil {
		respondStatus(w, http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	respondStatus(w, status)
	_, err = w.Write(respondWith)
	if err != nil {
		panic(err)
	}
}

func respondRedirect(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	respondStatus(w, http.StatusMovedPermanently)
}
