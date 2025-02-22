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