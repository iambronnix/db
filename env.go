package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	u "github.com/joho/godotenv"
)

func Config() (*sql.DB, error) {

	if loadErr := u.Load(); loadErr != nil {
		log.Fatal(loadErr)
	}

	dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, sqlErr := sql.Open("postgres", dbCreds)
	defer func() {
		for {
		pingErr := db.Ping()
			if  pingErr != nil { //something to fix here :either recover or scrap the whole infinite loop
				log.Fatal("....Connection has dropped.....")
				db.Close()

			}else{
				break
			}
			//fmt.Println("Connection is established...Good to go!!!")...remove comments in this line to debug connection issues
		
		}
	}()
	if sqlErr != nil {
		panic(sqlErr)
	}
	return db, nil

}
