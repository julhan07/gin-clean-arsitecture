package main

import (
	"clean-code/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()

	err := godotenv.Load()

	if err != nil {
		fmt.Println("env error", err)
		panic(err)
	}

	DB := database.IniDB()
	runApp(r, DB)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
