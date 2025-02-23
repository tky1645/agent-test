package plant

import (
	"DDD/entities"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// application service
type applicationService struct {
	repository IPlantRepository
}
type IPlantRepository interface {
	create(entities.Plant) error
	save(entities.Plant) error
	findByID(int) (entities.Plant, error)
}

type payloadPost struct {
	Name string `json:"name"`
}
type paramPatch struct {
	id int
}


func HandlerPOST(c *gin.Context) {
	// リポジトリの作成
	r := newRepo()

	param, err := fetchPost(c)
	if err !=nil{
		c.JSON(400, err)
		return
	}
	// バリデーション
	if err := validatePost(param); err != nil {
		c.JSON(400, err)
		return
	}
	// ルートエンティティの作成
	plant := entities.NewPlant(param.Name)
	// リポジトリ経由で保存
	if err := r.create(*plant); err != nil {
		c.JSON(500, err)
		return
	}
	
	fmt.Println(plant.Name)
	c.JSON(200, plant)
}

func HandlerPATCH(c *gin.Context) {
	r := newRepo()

	param, err := fetchPatch(c)
	if err !=nil{
		c.JSON(400, err)
		return
	}
	if err := validatePatch(param); err != nil {
		c.JSON(400, err)
		return
	}
	// ルートエンティティの再構築
	plant, err := r.findByID(param.id)
	if err != nil {
		c.JSON(404, err)
		return 
	}
	// 変更
	plant.UpdateWatering()

	// リポジトリ経由で保存
	if err := r.save(plant); err != nil {
		c.JSON(500, err)
		return
	}
	
	fmt.Println(plant.WateringDate)
	c.JSON(200, plant)
}

func fetchPost (c *gin.Context)(payloadPost,error){
	p := payloadPost{}
	if err :=c.ShouldBindJSON(&p); err !=nil{
		return p, err
	}
	return p, nil
}

func fetchPatch (c *gin.Context)(paramPatch,error){
	p :=c.Param("id");
	if  p ==""{
		return paramPatch{}, errors.New("id is empty")
	}

	id, err:= strconv.Atoi(p)
	if err != nil {
		return	paramPatch{}, err
	}

	return paramPatch{
		id: id,
	}, nil
}

func validatePost(p payloadPost) error {
	return nil
}
func validatePatch(p paramPatch) error {
	return nil
}