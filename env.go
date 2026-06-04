package db
import (
	"fmt"
	"log"
	"os"
	u "github.com/joho/godotenv"
  )
 func Config()(string, error){
	 if loadErr := u.Load();loadErr!=nil{
	 log.Fatal(loadErr)}
	 dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
	 os.Getenv("DB_USER"),
	 os.Getenv("DB_PASSWORD"),
	 os.Getenv("DB_HOST"),
	 os.Getenv("DB_PORT"),
	 os.Getenv("DB_NAME"),
 )
 return dbCreds, nil
 } 
