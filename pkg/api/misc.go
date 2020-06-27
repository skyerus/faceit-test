package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skyerus/faceit-test/pkg/customerror"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, 200, nil)
}

func getID(r *http.Request) (int, customerror.Error) {
	var ID int
	idStr, success := mux.Vars(r)["id"]
	if !success {
		return ID, customerror.NewGenericBadRequestError()
	}
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		return ID, customerror.NewBadRequestError("ID must be an integer")
	}

	return ID, nil
}
