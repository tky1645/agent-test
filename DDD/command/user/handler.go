package user

// application service
import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)


var (
	userRepository *UserRepository
	userService    *UserService
)

func InitHandlers(db *sql.DB) {
	userRepository = NewUserRepository(db)
	userService = NewUserService(*userRepository)
}

func HandlerGET(c *gin.Context) {
	user, err := userRepository.Create(1)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}

func HandlerPOST(c *gin.Context) {
	err := userService.Create(1, "postJohn")
	if err != nil {
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

func HandlerFETCH(c *gin.Context) {
	id := c.Param("id")
	user, err := userService.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func HandlerDELETE(c *gin.Context) {
	id := c.Param("id")
	err := userService.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User deleted successfully",
		"id":      id,
	})
}

func HandlerDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = userRepository.Delete(uint(idUint))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User deleted successfully",
		"id":      id,
	})
}
