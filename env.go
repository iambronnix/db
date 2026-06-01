package db
import (
	"errors"
	"fmt"
	"log"
	"os"
	u "github.com/joho/godotenv"
)
var (
	envERR = errors.New("error loading DB details")
)

func Config()(string, error){
	err := u.Load()
	if err != nil{
		log.Fatal(envERR)
	}//postgres://username:password@localhost:5432/dbname?sslmode=disable
	dbCreds := fmt.Sprintf("://%s:%s@%s:%d/%s?sslmode=disable")
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
)
return dbCreds, nil
}
// suspect there exist a bug in this package

