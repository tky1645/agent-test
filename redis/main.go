package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)
var cookieKey = "cookieKey"

func main() {
	r := gin.Default()

	redisRepo := &redisRepo{}
	redisRepo.New()
	loginCheckGroup := r.Group("/", chocklogin(redisRepo))
	{
		loginCheckGroup.GET("/", HandlerGet)
		
	}
	r.Run(":3000")

}

func HandlerGet(c *gin.Context){
	fmt.Println("call handler")
}

func chocklogin(repo repository)gin.HandlerFunc{
	return func (c *gin.Context)  {
		fmt.Println("call middleware")
		// sessionID取得
		var redisValue string
		var err error
		redisKey, sessionErr := getSessionID(c)
		if sessionErr != nil{
			fmt.Println("cannot find cookie")
			loginRequired(c, repo)
		}else if redisValue,err = repo.Get(redisKey); err ==  redis.Nil{
			// redis確認
			loginRequired(c, repo)
		}else if err != nil{
			fmt.Println(err.Error())
			c.Abort()
			return
		}else{
			fmt.Println("find id in redis")
		}
		fmt.Println(redisValue)

		// あればそのままNext()
		// ない時も諸々の処理を抜けてNext()
		c.Next()
		
	}
}

func login()error{
	fmt.Println("call login")
	return nil
}

func getSessionID(c  * gin.Context)(string, error){
	key, err := c.Cookie(cookieKey)
	return key, err
}

func loginRequired(c *gin.Context, repo repository){
	// ログイン
	fmt.Println("login Required")
	if err := login();err != nil{
		fmt.Println("fail to login")
		c.Abort()
		return
	}
	// セッション作成
	// セッションに登録する情報を生成
	user := user{
		Name: "starbacks",
		Nickname: "maccer",
	}
	// redisにセッションを登録
	uuid, _:= uuid.NewRandom()
	redisKey := uuid.String()
	repo.Create(c,redisKey, user)
	fmt.Println("redis data create")

	// cookieに保存
	fmt.Println(user)
	c.SetCookie(cookieKey, redisKey, 60,"/", "localhost", false, true)
}