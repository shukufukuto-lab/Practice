package main 

import "github.com/gin-gonic/gin"

func getList(C *gin.Context){
		c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main(){
	r := gin.Default()
	r.GET("/ping", getList)
	r.Run()
}