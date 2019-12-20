package main

import (
	"github.com/gin-contrib/cors"
	"go_learn/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.New()
	r.Use(cors.Default())

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(gin.Logger())

	r.Static("/static", "./static")

	routers.LoadRouters(r)

	//database.Init()
	//models.AutoMigrate()

	r.Run(":8081")
}
