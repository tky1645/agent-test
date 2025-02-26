package user

// application service
import (
	"github.com/gin-gonic/gin"
)

var(
 userRepository = NewUserRepository()
 userService = NewUserService(*userRepository)
)
func HandlerGET(c *gin.Context) {
	user,err  := userRepository.Create(1)
	if err !=nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}

func HandlerPOST(c *gin.Context) {
	err := userService.Create(1,"postJohn");
	if  err !=nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, err)
}

func HandlerPUT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}

	var name string
	if err := c.ShouldBindJSON(&name); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if name == "" {
		c.JSON(400, gin.H{"error": "name is empty"})
		return
	}
	

	err := userService.Update(id, name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User updated successfully",
		"id":      id,
	})
}
