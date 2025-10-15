package cmd

import (
	"admin-service/config"
	"admin-service/infra/db"
	"admin-service/repo"
	"admin-service/rest"
	"admin-service/rest/handlers/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.Getconfig()

	dbcon, err := db.Newconnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	err = db.Migrate(dbcon, "./migrations")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	//middlewares := middleware.NewMiddleware(cnf)
	userrepo := repo.Newuserrepo(dbcon)

	userHandler := user.NewHandler(cnf, userrepo)

	server := rest.Newserver(cnf, userHandler)

	server.Start()
}
