package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"redis_test/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)
var cookieKey = "cookieKey"

func Checklogin(repo util.Repository)gin.HandlerFunc{
	return func (c *gin.Context)  {
		fmt.Println("call middleware")
		// sessionID取得
		if !islogined(c, repo){
			loginRequired(c, repo)
		}

		// あればそのままNext()
		// ない時も諸々の処理を抜けてNext()
		c.Next()
		
	}
}

func islogined(c *gin.Context, repo util.Repository)bool{
	var redisValue string
	var err error
	redisKey, sessionErr := getSessionID(c)
	if sessionErr != nil{
		fmt.Println("cannot find cookie")
		return false
	}
	if redisValue,err = repo.Get(redisKey); err ==  redis.Nil{
		// redis確認
		return false
	}else if err != nil{
		fmt.Println(err.Error())
		c.Abort()
		return false
	}
	fmt.Println("find id in redis. add count")

	redisValue, _ = repo.Get(redisKey)
	var user util.User
	_ = json.Unmarshal([]byte(redisValue), &user)
	user.Count += 1
	repo.Create(c,redisKey,user) 

	return true
}

func login()error{
	fmt.Println("call login")
	return nil
}

func getSessionID(c  * gin.Context)(string, error){
	key, err := c.Cookie(cookieKey)
	return key, err
}

func loginRequired(c *gin.Context, repo util.Repository){
	// ログイン
	fmt.Println("login Required")
	if err := login();err != nil{
		fmt.Println("fail to login")
		c.Abort()
		return
	}
	// セッション作成
	// セッションに登録する情報を生成
	user := util.User{
		Name: "starbacks",
		Nickname: "maccer",
		Count: 1,
	}
	// redisにセッションを登録
	uuid, _:= uuid.NewRandom()
	redisKey := uuid.String()
	repo.Create(c,redisKey, user)
	fmt.Println("redis data create")

	// cookieに保存
	fmt.Println(user)
	
	// c.SetCookie(cookieKey, redisKey, 60,"/", "localhost", true, true)
	cookie := http.Cookie{
		Name:     cookieKey,
		Value:    redisKey,
		MaxAge:   60,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode, // SameSite 属性を設定
	}
	http.SetCookie(c.Writer, &cookie)
	
}