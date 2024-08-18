package routes

import (
	"github.com/RianAlvesAce/gestao_online/internal/controllers"
	"github.com/RianAlvesAce/gestao_online/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	main := r.Group("/api")
	{
		userRouter := main.Group("/user")
		{
			userRouter.GET("/login", controllers.UserLogin)
		}

		serviceRouter := main.Group("/service", middlewares.Auth())
		{
			serviceRouter.POST("/create", controllers.CreateService)
			serviceRouter.GET("/diagram", controllers.DiagramData)
		}

		pacientRouter := main.Group("/pacient", middlewares.Auth())
		{
			pacientRouter.POST("/create", controllers.CreatePacient)
			pacientRouter.PUT("/update", controllers.UpdatePacient)
		}

		scheduleRouter := main.Group("/schedule", middlewares.Auth())
		{
			scheduleRouter.POST("/create", controllers.CreateSchedule)
			scheduleRouter.GET("/diagram", controllers.GetSchedule)
		}

	}	
}