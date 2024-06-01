package dbconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var User string
var PostgresDriver string
var Host string
var Port string
var Password string
var DbName string
var TableName string
var DataSourceName string

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("ERROR ENV VAR = ", err.Error())
	}
    User = os.Getenv("POSTGRES_USER")
    PostgresDriver = "postgres"
    Host = os.Getenv("POSTGRES_HOST")
    Port = os.Getenv("DB_PORT")
    Password = os.Getenv("DB_PASSWORD")
    DbName = "spyTroth"
    TableName = "messages"

    DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
}



