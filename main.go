package main

import (
 "fmt"

 "github.com/gin-gonic/gin"
)

func main() {
 r := gin.Default()

 r.POST("/", func(c *gin.Context) {
   fmt.Println("resp")
   buf := make([]byte, 2048)
   n, _ := c.Request.Body.Read(buf)
   b := string(buf[0:n])
   fmt.Println(b)
 })

 r.Run(":8080")
}
