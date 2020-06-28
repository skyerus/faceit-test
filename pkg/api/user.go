package api

import (
	"encoding/json"
	"net/http"

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
	ID, customErr := getID(r)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	u, found := router.c.GetUsersCache().GetUser(ID)
	if found {
		respondJSON(w, http.StatusOK, u)
		return
	}
	userRepo := userrepo.NewMysqlUserRepository(router.conn)
	userService := userservice.NewUserService(userRepo)
	u, customErr = userService.Get(ID)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	go router.c.GetUsersCache().AddUser(u)

	respondJSON(w, http.StatusOK, u)
}

func (router *router) deleteUser(w http.ResponseWriter, r *http.Request) {
	ID, customErr := getID(r)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	userRepo := userrepo.NewMysqlUserRepository(router.conn)
	userService := userservice.NewUserService(userRepo)
	customErr = userService.Delete(ID)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	router.c.GetUsersCache().DeleteUser(ID)

	respondJSON(w, http.StatusNoContent, nil)
}

func (router *router) getUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var filter user.Filter
	filter = make(map[string]string)
	filter["nickname"] = query.Get("nickname")
	filter["country"] = query.Get("country")

	userRepo := userrepo.NewMysqlUserRepository(router.conn)
	userService := userservice.NewUserService(userRepo)
	users, customErr := userService.GetAll(filter)
	if customErr != nil {
		handleError(w, customErr)
		return
	}

	respondJSON(w, http.StatusOK, users)
}

func (router *router) updateUser(w http.ResponseWriter, r *http.Request) {
	ID, customErr := getID(r)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		respondBadRequest(w)
		return
	}
	u.ID = ID

	userRepo := userrepo.NewMysqlUserRepository(router.conn)
	userService := userservice.NewUserService(userRepo)
	customErr = userService.Update(u)
	if customErr != nil {
		handleError(w, customErr)
		return
	}
	router.c.GetUsersCache().AddUser(u)

	respondJSON(w, http.StatusOK, u)
}
