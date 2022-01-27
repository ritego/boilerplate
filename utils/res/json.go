package res

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, r Response, status int64) {
	if status == 0 {
		status = http.StatusOK
	}
	w.WriteHeader(int(status))
	json.NewEncoder(w).Encode(r)
}
