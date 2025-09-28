package plant

import (
	"DDD/entities"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type IPlantRepository interface {
	create(entities.Plant) error
	save(entities.Plant) error
	findByID(string) (entities.Plant, error)
	FindAll(limit int, offset int) ([]entities.Plant, error)
	FindWateringRecordsByPlantID(plantID string) ([]entities.WateringRecord, error)
	CreateWateringRecord(entities.WateringRecord) error
}

type PlantRepository interface {
}

type payloadPost struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
}
type paramPatch struct {
	id string
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
	plant := entities.NewPlant(param.Name, entities.WithDescription(param.Description), entities.WithImageURL(param.ImageURL))
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

	return paramPatch{
		id: p,
	}, nil
}

func validatePost(p payloadPost) error {
	return nil
}
func validatePatch(p paramPatch) error {
	return nil
}

type PlantListRequest struct {
	Limit  int `json:"limit" binding:"required"`
	Offset int `json:"offset" binding:"required"`
}

type WateringRecordRequest struct {
	Notes *string `json:"notes"`
}

func HandlerGETPlants(c *gin.Context) {
	var req PlantListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := newRepo()
	plants, err := repo.FindAll(req.Limit, req.Offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, plants)
}

func HandlerGETWateringHistory(c *gin.Context) {
	plantID := c.Param("plantId")
	if plantID == "" {
		c.JSON(400, gin.H{"error": "plant_id is required"})
		return
	}

	repo := newRepo()
	records, err := repo.FindWateringRecordsByPlantID(plantID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, records)
}

func generateUUID() string {
	return uuid.New().String()
}

func HandlerPOSTWatering(c *gin.Context) {
	plantID := c.Param("id")
	if plantID == "" {
		c.JSON(400, gin.H{"error": "plant_id is required"})
		return
	}

	var req WateringRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := newRepo()
	
	recordID := generateUUID()
	
	record := entities.WateringRecord{
		ID:        recordID,
		PlantID:   plantID,
		WateredAt: time.Now(),
		Notes:     req.Notes,
		CreatedAt: time.Now(),
	}

	if err := repo.CreateWateringRecord(record); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	plant, err := repo.findByID(plantID)
	if err != nil {
		c.JSON(404, gin.H{"error": "plant not found"})
		return
	}
	
	plant.UpdateWatering()
	if err := repo.save(plant); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, record)
}
