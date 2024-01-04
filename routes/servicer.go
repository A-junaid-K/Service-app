package routes

import (
	"service-api/controllers"

	"github.com/gin-gonic/gin"
)

func ServiceRouter(server *gin.Engine) {
	//grouping all routes with services with "r"
	r := server.Group("/servicer")

	r.POST("/signup", controllers.ServicerSignup)
	r.POST("/adddocuments/:id", controllers.AddDocuments)
	r.GET("/getdetails/:servicer_id", controllers.GetServicerDetails)
	r.POST("/login", controllers.ServicerLogin)
	r.GET("/getallbooking/:servicer_id", controllers.GetAllBookings)
	r.POST("/changestatus/:servicer_id", controllers.ChangeBookingStatus)
}
