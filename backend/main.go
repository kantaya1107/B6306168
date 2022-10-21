package main

import (
	"github.com/kantaya1107/sa-65-example/controller"

	"github.com/kantaya1107/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	//Leave Routes

	r.GET("/leaves", controller.ListLeave)

	r.GET("/leaves/:id", controller.GetLeave)

	r.POST("/leaves", controller.CreateLeave)

	r.GET("/doctors", controller.ListDoctor)
	r.GET("/types", controller.ListType)
	r.GET("/evidences", controller.ListEvidence)

	//  r.PATCH("/users", controller.UpdateUser)

	//  r.DELETE("/users/:id", controller.DeleteUser)

	//Run the server

	r.Run()

}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
