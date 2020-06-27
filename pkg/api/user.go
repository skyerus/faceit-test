package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skyerus/faceit-test/pkg/user"
	"github.com/skyerus/faceit-test/pkg/user/userrepo"
	"github.com/skyerus/faceit-test/pkg/user/userservice"
)

func (router *router) createUser(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		respondBadRequest(w)
		return
	}
	userRepo := userrepo.NewMysqlUserRepository(router.conn)
	userService := userservice.NewUserService(userRepo)
	customError := userService.Create(&u)
	if customError != nil {
		handleError(w, customError)
		return
	}

	respondJSON(w, http.StatusCreated, u)
}

func (router *router) getUser(w http.ResponseWriter, r *http.Request) {
	idStr, success := mux.Vars(r)["id"]
	if !success {
		respondBadRequest(w)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondBadRequest(w)
		return
	}
	userRepo := userrepo.NewMysqlUserRepository(router.conn)
	userService := userservice.NewUserService(userRepo)
	u, customErr := userService.Get(id)
	if customErr != nil {
		handleError(w, customErr)
		return
	}

	respondJSON(w, http.StatusOK, u)
}
