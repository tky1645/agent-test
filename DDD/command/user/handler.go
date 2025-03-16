package user

// application service
import (
	"database/sql"
	"fmt"

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

// エラーレスポンスの共通化
func sendErrorResponse(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

// 成功レスポンスの共通化
func sendSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func HandlerGET(c *gin.Context) {
	users, err := userRepository.GetAll()
	if err != nil {
		sendErrorResponse(c, 500, err)
		return
	}
	sendSuccessResponse(c, users)
}

func HandlerPOST(c *gin.Context) {
	// POSTではリクエストボディからデータを受け取るように修正
	var request struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		sendErrorResponse(c, 400, err)
		return
	}
	
	err := userService.Create(1, request.Name)
	if err != nil {
		sendErrorResponse(c, 500, err)
		return
	}
	sendSuccessResponse(c, gin.H{"message": "User created successfully"})
}

func HandlerPUT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendErrorResponse(c, 400, fmt.Errorf("id is empty"))
		return
	}

	var request struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		sendErrorResponse(c, 400, err)
		return
	}

	err := userService.Update(id, request.Name)
	if err != nil {
		sendErrorResponse(c, 500, err)
		return
	}
	sendSuccessResponse(c, gin.H{
		"message": "User updated successfully",
		"id":     id,
	})
}

// HandlerFETCHとHandlerDELETEを統合し、命名規則を統一
func HandlerGetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendErrorResponse(c, 400, fmt.Errorf("id is empty"))
		return
	}
	
	user, err := userService.GetByID(id)
	if err != nil {
		sendErrorResponse(c, 500, err)
		return
	}
	sendSuccessResponse(c, user)
}

func HandlerDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendErrorResponse(c, 400, fmt.Errorf("id is empty"))
		return
	}

	err := userService.Delete(id)
	if err != nil {
		sendErrorResponse(c, 500, err)
		return
	}
	sendSuccessResponse(c, gin.H{
		"message": "User deleted successfully",
		"id":     id,
	})
}
