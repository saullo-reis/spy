package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	dbconfig "spy.com/database/config"
	"spy.com/database/migrations"
)

type message struct {
	Message string `json:"message"`
}

func main(){
	migrations.CreateTables()

	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)
	if err != nil {
		panic("Error connecting in database" + err.Error())
	}
	fmt.Println("Connection with database sucess")
	defer db.Close()

	r := gin.New()
    r.Use(gin.Recovery())
	
	r.POST("/send-message", func(c * gin.Context){
		var message message
		err := c.ShouldBindJSON(&message)
		if err != nil {
			fmt.Println("ERROR READING JSON")
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "ERROR READING JSON",
			})
			return
		}

		_, err = db.Exec("INSERT INTO MESSAGES (MESSAGE) VALUES ($1)", message.Message)
		if err != nil{
			fmt.Println("ERROR INSERTING MESSAGE")
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"message": "ERROR INTERNAL IN INSERT VALUE",
			} )
			return
		} 
		fmt.Println("MESSAGE SENT "+ message.Message)

	})
	r.Run()
}