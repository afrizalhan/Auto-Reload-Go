package main

import (
	"net/http"
	"auto-reload-go/service"
	"github.com/gin-gonic/gin"
)

func main(){
	go service.Reloader()
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.GET("/", index)
	r.Static("/assets", "template/assets")
	r.Run(":8080")
}

func index(c *gin.Context){
	data := service.ReadJSON()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"water":      data["statusWater"],
		"wind":       data["statusWind"],
		"waterValue": data["water"],
		"windValue":  data["wind"],
	})
}