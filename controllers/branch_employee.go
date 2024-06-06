package controllers

import (
	"InventoryManagement/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (b BranchController) CreateBranchController(c *gin.Context) {
	var branch models.Branch

	if err := c.ShouldBindJSON(&branch); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	res, err := branch.Create()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data":res})
}

func (b BranchController) UpdateBranchController(c *gin.Context) {
	var branchUpdated models.Branch
	var branch models.Branch
	pk_str := c.Param("id")
	pk, err := strconv.Atoi(pk_str)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	if err := c.ShouldBindJSON(&branchUpdated); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	res, err := branch.Update(pk, branchUpdated)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data":res})
}

func (b BranchController) DeleteBranchController(c *gin.Context) {
	var branch models.Branch
	pk_str := c.Param("id")
	pk, err := strconv.Atoi(pk_str)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	if err := c.ShouldBindJSON(&branch); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	err = branch.Delete(pk)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message":"Entry deleted"})
}

func (b BranchController) GetBranchByOneController(c *gin.Context) {
	var branch models.Branch
	pk_str := c.Param("id")
	fmt.Println("id",pk_str)
	pk, err := strconv.Atoi(pk_str)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	if err := c.ShouldBindJSON(&branch); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	res, err := branch.FindById(pk)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data":res})
}

func (b BranchController) GetBranchController(c *gin.Context) {
	data, err := models.GetAllBranches()
	if err != nil {
		c.JSON(400, gin.H{"error": "Unable to Fetch data"})
		return
	}
	c.JSON(200, gin.H{"data": data})
}

func (t TeamController) CreateTeamController(c *gin.Context) {
	var team models.Team

	if err := c.ShouldBindJSON(&team); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	res, err := team.Create()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data":res})
}

func (t ParkingController) CreateParkingController(c *gin.Context) {
	var parking models.ParkingSpot

	if err := c.ShouldBindJSON(&parking); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	res, err := parking.Create()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data":res})
}
