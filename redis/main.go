package main

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)
var cookieKey = "key"

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
	return nil
}

func getSessionID(c  * gin.Context)(string, error){
	key, err := c.Cookie(cookieKey)
	return key, err
}

func loginRequired(c *gin.Context, repo repository){
	// ログイン
	fmt.Println("cannot find redis data")
	if err := login();err != nil{
		fmt.Println("fail to login")
		c.Abort()
		return
	}
	// セッション作成
	// セッションに登録する情報を生成
	user := user{
		name: "starbacks",
	}
	// redisにセッションを登録
	b := make([]byte, 64)
	redisKey:= base64.URLEncoding.EncodeToString(b)
	repo.Create(redisKey, user)
	fmt.Println("redis data create")

	// cookieに保存
	c.SetCookie(cookieKey, redisKey, 60,"/", "localhost", false, true)
}