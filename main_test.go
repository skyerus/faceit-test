package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/skyerus/faceit-test/pkg/api"
	"github.com/skyerus/faceit-test/pkg/db"
	"github.com/skyerus/faceit-test/pkg/env"
	"github.com/skyerus/faceit-test/pkg/user"
)

var a api.App
var conn *sql.DB
var testUser user.User = user.User{
	ID:        1,
	FirstName: "John",
	LastName:  "Appleseed",
	Email:     "john.appleseed@gmail.com",
	Nickname:  "jappleseed",
	Country:   "UK",
}

func TestMain(m *testing.M) {
	var err error
	env.SetEnv()
	a = api.App{}
	conn, err = db.OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = db.CreateUserTable(conn)
	if err != nil {
		log.Fatal(err)
	}
	a.Initialize(conn)
	code := m.Run()
	err = db.ClearUserTable(conn)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestGetNonExistentUser(t *testing.T) {
	db.ClearUserTable(conn)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	assertResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	assertErrorMessage(t, "No user exists with id 1", m["message"])
}

func TestCreateUser(t *testing.T) {
	db.ClearUserTable(conn)
	response := createUser(testUser)
	assertResponseCode(t, http.StatusCreated, response.Code)

	var u user.User
	err := json.Unmarshal(response.Body.Bytes(), &u)
	if err != nil {
		t.Fatalf("Unmarshal error")
	}
	assertUsers(t, testUser, u)
}

func TestGetUser(t *testing.T) {
	db.ClearUserTable(conn)
	createUser(testUser)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	assertResponseCode(t, http.StatusOK, response.Code)
	var u user.User
	err := json.Unmarshal(response.Body.Bytes(), &u)
	if err != nil {
		t.Fatalf("Unmarshal error")
	}
	assertUsers(t, testUser, u)
}

func TestDeleteUser(t *testing.T) {
	db.ClearUserTable(conn)
	createUser(testUser)
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	response := executeRequest(req)
	assertResponseCode(t, http.StatusNoContent, response.Code)
	req, _ = http.NewRequest("GET", "/users/1", nil)
	response = executeRequest(req)
	assertResponseCode(t, http.StatusNotFound, response.Code)
	// Test idempotency
	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	response = executeRequest(req)
	assertResponseCode(t, http.StatusNoContent, response.Code)
}

func TestGetUsers(t *testing.T) {
	db.ClearUserTable(conn)
	createUser(testUser)
	otherUser := testUser
	otherUser.Nickname = "anna"
	otherUser.Email = "anna@gmail.com"
	otherUser.Country = "UK"
	createUser(otherUser)
	otherUser.Nickname = "josh"
	otherUser.Email = "josh@gmail.com"
	otherUser.Country = "France"
	createUser(otherUser)

	var filter user.Filter
	filter = make(map[string]string)
	users := getUsers(t, filter)
	if len(users) != 3 {
		t.Errorf("Expected length of result to be '%v'. Got '%v'", 3, len(users))
	}

	filter = make(map[string]string)
	filter["country"] = "UK"
	users = getUsers(t, filter)
	if len(users) != 2 {
		t.Errorf("Expected length of result to be '%v'. Got '%v'", 2, len(users))
	}

	filter = make(map[string]string)
	filter["nickname"] = "j"
	users = getUsers(t, filter)
	if len(users) != 2 {
		t.Errorf("Expected length of result to be '%v'. Got '%v'", 2, len(users))
	}
	filter["country"] = "UK"
	users = getUsers(t, filter)
	if len(users) != 1 {
		t.Errorf("Expected length of result to be '%v'. Got '%v'", 1, len(users))
	}
}

func TestUpdateUser(t *testing.T) {
	db.ClearUserTable(conn)
	createUser(testUser)
	updatedUser := testUser
	updatedUser.Nickname = "johnnyapple"
	byteData, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(byteData))
	response := executeRequest(req)
	assertResponseCode(t, http.StatusOK, response.Code)
	req, _ = http.NewRequest("GET", "/users/1", nil)
	response = executeRequest(req)
	var u user.User
	err := json.Unmarshal(response.Body.Bytes(), &u)
	if err != nil {
		t.Fatalf("Unmarshal error")
	}
	assertUsers(t, updatedUser, u)
}

func getUsers(t *testing.T, filter user.Filter) []user.User {
	req, _ := http.NewRequest("GET", "/users", nil)
	q := req.URL.Query()
	for k, v := range filter {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	response := executeRequest(req)
	assertResponseCode(t, http.StatusOK, response.Code)
	var users []user.User
	err := json.Unmarshal(response.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("Unmarshal error")
	}

	return users
}

func assertUsers(t *testing.T, expected user.User, actual user.User) {
	if actual.FirstName != expected.FirstName {
		t.Errorf("Expected first name to be '%v'. Got '%v'", expected.FirstName, actual.FirstName)
	}
	if actual.LastName != expected.LastName {
		t.Errorf("Expected last name to be '%v'. Got '%v'", expected.LastName, actual.LastName)
	}
	if actual.Email != expected.Email {
		t.Errorf("Expected email to be '%v'. Got '%v'", expected.Email, actual.Email)
	}
	if actual.Nickname != expected.Nickname {
		t.Errorf("Expected nickname to be '%v'. Got '%v'", expected.Nickname, actual.Nickname)
	}
	if actual.Country != expected.Country {
		t.Errorf("Expected country to be '%v'. Got '%v'", expected.Country, actual.Country)
	}
	if actual.ID != expected.ID {
		t.Errorf("Expected user ID to be '%v'. Got '%v'", expected.ID, actual.ID)
	}
}

func createUser(u user.User) *httptest.ResponseRecorder {
	byteData, _ := json.Marshal(u)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(byteData))
	req.Header.Set("Content-Type", "application/json")
	return executeRequest(req)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func assertResponseCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func assertErrorMessage(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected the 'message' key of the response to be set to '%s'. Got '%s'", expected, actual)
	}
}
