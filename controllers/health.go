package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




func (h HealthController) Status(c *gin.Context){
	c.String(http.StatusOK, "Working !")
}

func (b BranchController) CreateController(c *gin.Context){
	
}