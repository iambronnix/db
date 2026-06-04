package db

import (
<<<<<<< HEAD
	"database/sql"
=======
>>>>>>> refs/remotes/origin/main
	"fmt"
	"log"
	"os"

	u "github.com/joho/godotenv"
)
 func Config()(*sql.DB, error){
 
	 if loadErr := u.Load();loadErr!=nil{
	 log.Fatal(loadErr)}
     
	 dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
	 os.Getenv("DB_USER"),
	 os.Getenv("DB_PASSWORD"),
	 os.Getenv("DB_HOST"),
	 os.Getenv("DB_PORT"),
	 os.Getenv("DB_NAME"),
 )
		db, sqlErr := sql.Open("postgres",dbCreds)
		defer func(){
			for {
			if pingErr := db.Ping();pingErr!=nil{
				fmt.Println("....Connection has dropped.....")	
                db.Close()			
			}else{
				fmt.Println("Connection is established...Good to go!!!")
				break
			}
		}
		}()
		if sqlErr!=nil{
	  panic(sqlErr)
		}
		return db, nil		
 
 } 