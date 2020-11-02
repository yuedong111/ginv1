package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	r.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "defaultvalue")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id:%s; page:%s; name:%s; message:%s", id, page, name, message)
	})
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context){
		file, _ := c.FormFile("file")
		fmt.Printf(file.Filename)
		if err := c.SaveUploadedFile(file, "D:/code/ginv1/test.txt"); err != nil{
			fmt.Println(err)
		}
		c.String(200, fmt.Sprintf("%s uploaded", file.Filename))

	})
	r.Run(":8080")
}
