package main 

import "github.com/gin-gonic/gin"

func getting(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func posting(c *gin.Context) {
    /*
       DB操作など
    */
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}


func main(){
	r := gin.Default()
	r.GET("/get", getting)
	r.POST("/post", posting)
	r.Run()
}