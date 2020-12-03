package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/mygreeting",addCors, func(c *gin.Context){
		c.JSON(200, gin.H{
			"owner": "Martín Morales",
			"greeting": "Hola, Soy martin y estoy en el equipo de reporting. ",
			"repository": "https://github.com/mamoralesMELI/go-hands-on",
		})
	})
	r.Run()
}

func addCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Request-Method")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	c.Header("Allow", "GET, POST, OPTIONS, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}