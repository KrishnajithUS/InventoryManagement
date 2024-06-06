package controllers

import (
	"InventoryManagement/models"
	"InventoryManagement/utils/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r UserController) CurrentUser(c *gin.Context){

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}

	u,err := user.GetUserById(int(user_id))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":u})
}

func (r UserController) RegisterUserController(c *gin.Context){
	var userInput RegisterInput
	if err := c.ShouldBindBodyWithJSON(&userInput); err != nil{
		c.JSON(400, gin.H{"error":err})
		return 
	}
	user := models.User{}
	user.Username = userInput.Username
	user.Password = userInput.Password
	_, err := user.CreateUser()
	fmt.Println("error",err)
	if err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
	}
	c.JSON(200, gin.H{"message":"Register controller"})
}


func (r UserController) LoginUserController(c *gin.Context){
	var userInput LoginInput
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&userInput); err != nil{
		c.JSON(400, gin.H{"error":err})
		return 
	}
	username := userInput.Username
	password := userInput.Password
	token, err := user.LoginUser(username, password)
	if err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}

func (r EmployeeController) CreateEmployeeController(c *gin.Context){
	var userInput EmployeeCreateInput
	if err := c.ShouldBindBodyWithJSON(&userInput); err != nil{
		c.JSON(400, gin.H{"error":err})
		return 
	}
	user := models.Employee{}
	user.Username = userInput.Username
	user.Password = userInput.Password
	user.TeamId = userInput.TeamId
	user.PersonalInfo = userInput.PersonalInfo
	_, err := user.CreateUser()
	fmt.Println("error",err)
	if err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
	}
	c.JSON(200, gin.H{"message":"Register controller"})
}


func (r EmployeeController) LoginEmployeeController(c *gin.Context){
	var userInput EmployeeLoginInput
	var user models.Employee
	if err := c.ShouldBindBodyWithJSON(&userInput); err != nil{
		c.JSON(400, gin.H{"error":err})
		return 
	}
	username := userInput.Username
	password := userInput.Password
	token, err := user.LoginUser(username, password)
	if err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}