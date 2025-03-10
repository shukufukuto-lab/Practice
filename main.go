package main 

import "github.com/gin-gonic/gin"
import "net/http"

func getting(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getting2(c *gin.Context){
	c.JSON(200, gin.H{
		"messageaaaaa": "pong",
	})
}


func posting(c *gin.Context) {
    /*
       DB操作など
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
		    */
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}


func main(){
	r := gin.Default()
	r.GET("/get", getting)
	r.GET("/get2", getting2)
	r.POST("/post", posting)
	r.Run()
}