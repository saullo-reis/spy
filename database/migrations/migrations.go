package migrations

import (
	_ "github.com/lib/pq"
	dbconfig "spy.com/database/config"
	"database/sql"
	"fmt"
)

func CreateTables(){

	queryToCreateTables := `
		CREATE TABLE IF NOT EXISTS MESSAGES(
			MESSAGE TEXT,
			UPDATED_DATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CREATED_DATE TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)
	if err != nil {
		panic("Error connecting in database" + err.Error())
	}
	fmt.Println("Connection with database sucess")
	defer db.Close()

	_, err = db.Exec(queryToCreateTables)
	if err != nil{
		panic("Error in migrations: "+ err.Error())
	}
}