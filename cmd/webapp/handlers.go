package webapp

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"net/http"
)

type Entity struct {
	Code string
	Url string
}

func shorten(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	r.ParseForm()

	code := r.Form.Get("code")
	url := r.Form.Get("url")

	if url == "" {
		respondJSON(w, http.StatusBadRequest, "Missing url")
	}

	if code == "" {
		respondJSON(w, http.StatusBadRequest, "Missing Code")
	}

	client, err := datastore.NewClient(ctx, "dbut-0")
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
	}

	k := datastore.IncompleteKey("Url", nil)

	e := new(Entity)

	e.Code = code
	e.Url = url

	_, err = client.Put(ctx, k, e)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
	}

	respondJSON(w, http.StatusOK, "Ok")

}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	respondWith, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(respondWith)
	if err != nil {
		panic(err)
	}
}
