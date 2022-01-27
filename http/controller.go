package http

import (
	"net/http"
)

type ctrl struct{}

func (ctrl *ctrl) home(rw http.ResponseWriter, r *http.Request) {

}

func (ctrl *ctrl) health(rw http.ResponseWriter, r *http.Request) {
	// ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	// defer cancel()

	// err := s.db.PingContext(ctx)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("db down: %v", err), http.StatusFailedDependency)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// return
}

func (ctrl *ctrl) settlement(rw http.ResponseWriter, r *http.Request) {

}

func (ctrl *ctrl) payout(rw http.ResponseWriter, r *http.Request) {

}

var controller = &ctrl{}
