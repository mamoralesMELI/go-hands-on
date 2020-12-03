package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamoralesMELI/go-hands-on/rest"
)

type Greeting struct {
	greeting string `json:"greeting"`
	owner string `json:"owner"`
	repository string `json:"repository"`
}

func main() {
	r := gin.Default()
	r.GET("/mygreeting",addCors, func(c *gin.Context){
		r := rest.New()
		g := new(Greeting)
		r.Get("https://fanaur-workshop-go-hands-on.herokuapp.com/fanaur",nil, g)

		c.JSON(200, gin.H{
			"owner": "Martín Morales",
			"greeting": g.greeting,
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