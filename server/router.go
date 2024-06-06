package server

import (
	"InventoryManagement/controllers"
	"InventoryManagement/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	health := new(controllers.HealthController)
	branch := new(controllers.BranchController)
	team := new(controllers.TeamController)
	parking := new(controllers.ParkingController)
	user := new(controllers.UserController)
	employee := new(controllers.EmployeeController)
	// Public
	router.POST("register/", user.RegisterUserController)
	router.POST("login/", user.LoginUserController)
	router.GET("healthz/", health.Status)
	router.POST("login-employee/", employee.LoginEmployeeController)

	router.Use(middlewares.JwtAuthMiddleware())
	// Private
	router.POST("create-team/", team.CreateTeamController)
	router.POST("create-parking-spot/", parking.CreateParkingController)
	router.GET("branches/", branch.GetBranchController)
	router.POST("branches/", branch.CreateBranchController)
	router.GET("branches/:id", branch.GetBranchByOneController)
	router.PUT("branches/:id", branch.UpdateBranchController)
	router.DELETE("branches/:id", branch.DeleteBranchController)
	router.POST("create-employee/", employee.CreateEmployeeController)

	return router
}
