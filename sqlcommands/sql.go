package sqllayer

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host	 = "localhost"
	port	 = 5432
	user	 = "postgres"
	sbCustomerDb = "SB_Customers"
)

func ConnectToDB(password  string){
	sqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, sbCustomerDb)
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	command := "\\dt"
	rows, err := db.Query(command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}