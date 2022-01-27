package req

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Body(r *http.Request, v interface{}) {
	b, _ := io.ReadAll(r.Body)
	json.Unmarshal(b, v)
}

func Param(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}

func Query(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
