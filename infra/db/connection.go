package db

import (
	"admin-service/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Newconnection(cnf *config.DBcofig) (*sqlx.DB, error) {
	dbsource := GetConnectionString(cnf)
	dbcon, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbcon, nil
}

func GetConnectionString(cnf *config.DBcofig) string {
	connstring := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s ",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)

	if !cnf.EnableSSLMODE {
		connstring += " sslmode=disable"
	}
	return connstring
}
