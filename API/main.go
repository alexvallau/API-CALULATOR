package main

import (
	"github.com/gin-gonic/gin"

	controllers "API/controllers"
)

func main() {
	controllers.InitDataBase()

	r := gin.Default()

	r.GET("/books", controllers.FindBooks)
	r.GET("/add/:a/:b", controllers.Add)
	r.GET("/multiply/:a/:b", controllers.Multiply)
	r.GET("/minus/:a/:b", controllers.Minus)
	r.Run()

}
