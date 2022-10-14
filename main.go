package main

import (
	"fmt"
	//sqllayer "store-sense/WS-to-DB/app/sqlcommands"
	"errors"
	"os"
	sqllayer "store-sense/WS-to-DB/app/sqlcommands"

	"github.com/joho/godotenv"
)

func getDBPassword() (response string, errObj error){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", errors.New("Error loading .env file")
	}
	dbpassword := os.Getenv("password")
	if dbpassword == ""{
		fmt.Println("No password found in .env file with key 'password'")
		return "", errors.New("No password found in .env file with key 'password'")
	}
	return dbpassword, nil
}

func main()  {

	 dbPassword, err := getDBPassword()
	 if err != nil {
		 fmt.Println("Error getting DB password")
			fmt.Printf("panic: %v\n", err)
			panic(err)
		}
		sqllayer.ConnectToDB(dbPassword)
}