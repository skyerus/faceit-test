package userrepo

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/skyerus/faceit-test/pkg/customerror"
	"github.com/skyerus/faceit-test/pkg/user"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMysqlUserRepository ...
func NewMysqlUserRepository(Conn *sql.DB) user.Repository {
	return &mysqlUserRepository{Conn}
}

func (ur mysqlUserRepository) Create(u *user.User) customerror.Error {
	stmtIns, err := ur.Conn.Prepare("INSERT INTO user (first_name, last_name, nickname, email, country) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(u.FirstName, u.LastName, u.Nickname, u.Email, u.Country)
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return customerror.NewUnprocessableEntity("User already exists")
		}
	}
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	u.ID = int(id64)

	return nil
}

func (ur mysqlUserRepository) Get(ID int) (user.User, customerror.Error) {
	var u user.User
	results, err := ur.Conn.Query("SELECT user.id, user.first_name, user.last_name, user.nickname, user.email, user.country FROM user WHERE id = ?", ID)
	if err != nil {
		return u, customerror.NewGenericHTTPError(err)
	}
	defer results.Close()
	res := results.Next()
	if !res {
		return u, customerror.NewNotFoundError("No user exists with id " + strconv.Itoa(ID))
	}
	err = results.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Nickname, &u.Email, &u.Country)
	if err != nil {
		return u, customerror.NewGenericHTTPError(err)
	}

	return u, nil
}

func (ur mysqlUserRepository) Delete(ID int) customerror.Error {
	stmtIns, err := ur.Conn.Prepare("DELETE FROM `user` WHERE id = ?")
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(ID)
	if err != nil {
		return customerror.NewGenericHTTPError(err)
	}

	return nil
}

func (ur mysqlUserRepository) GetAll(f user.Filter) ([]user.User, customerror.Error) {
	users := make([]user.User, 0)
	var qb strings.Builder
	qb.WriteString("SELECT * FROM user")
	appendFilterToQuery(&qb, f)

	results, err := ur.Conn.Query(qb.String())
	if err != nil {
		return users, customerror.NewGenericHTTPError(err)
	}
	defer results.Close()
	for results.Next() {
		var u user.User
		err = results.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Nickname, &u.Email, &u.Country)
		if err != nil {
			return users, customerror.NewGenericHTTPError(err)
		}
		users = append(users, u)
	}

	return users, nil
}

func appendFilterToQuery(qb *strings.Builder, f user.Filter) {
	var filterBeenApplied bool
	for k, v := range f {
		if v == "" {
			continue
		}
		if !filterBeenApplied {
			qb.WriteString(" WHERE ")
			filterBeenApplied = true
		} else {
			qb.WriteString(" AND ")
		}
		qb.WriteString(k)
		qb.WriteString(" LIKE '%")
		qb.WriteString(v)
		qb.WriteString("%'")
	}
}
