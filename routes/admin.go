package routes

import (
	"service-api/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(server *gin.Engine) {
	r := server.Group("/admin")
	r.GET("/pendingservicer", controllers.GetPendingServicer)
	r.GET("/acceptedservicer", controllers.GetAcceptedServicer)
	r.GET("/rejectedservicer", controllers.GetRejectedServicer)
	r.PATCH("/acceptservicer/:servicer_id", controllers.AcceptServicer)
	r.PATCH("/rejectservicer/:servicer_id", controllers.RejectServicer)
}
