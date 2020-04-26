package webapp

import (
	"encoding/json"
	"github.com/dbut2/shortener/pkg/shortener"
	"github.com/go-chi/chi"
	"net/http"
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
		respondJSON(w, http.StatusBadRequest, "Missing url")
	}
	if code == "" {
		respondJSON(w, http.StatusBadRequest, "Missing Code")
	}

	err := shortener.Shorten(url, code)
	if err != nil {
		respondStatus(w, http.StatusInternalServerError)
	}

	respondJSON(w, http.StatusOK, "Ok")

}

func lengthen(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := chi.URLParam(r, "code")

	if !shortener.CodeExists(code) {
		respondStatus(w, http.StatusNotFound)
	}

	url, err := shortener.Lengthen(code)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
	}

	respondRedirect(w, url)
}

func respondStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
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
