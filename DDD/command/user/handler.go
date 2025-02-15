package user

// application service
import (
	"DDD/entities"

	"github.com/gin-gonic/gin"
)

var userRepository = NewUserRepository()

func HandlerGET(c *gin.Context) {
	user  := userRepository.Create(1)
	c.JSON(200, user)
}

func HandlerPOST(c *gin.Context) {
	user  := entities.NewUser(1, "postJohn")
	if err := userRepository.Save(user); err !=nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}