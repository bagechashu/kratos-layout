package swagger

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewSwaggerUIRouter() *mux.Router {
	router := mux.NewRouter()
	handler := http.StripPrefix("/swagger", http.FileServer(http.Dir("../../swagger/ui")))
	router.PathPrefix("/swagger").Handler(handler)
	return router
}
