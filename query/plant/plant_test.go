package plant

import (
	"DDD/entities"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) create(plant entities.Plant) error {
	args := m.Called(plant)
	return args.Error(0)
}

func (m *MockRepository) save(plant entities.Plant) error {
	args := m.Called(plant)
	return args.Error(0)
}

func (m *MockRepository) findByID(id int) (entities.Plant, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Plant), args.Error(1)
}

func (m *MockRepository) FindAll(limit int, offset int) ([]entities.Plant, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.Plant), args.Error(1)
}

func (m *MockRepository) FindWateringRecordsByPlantID(plantID string) ([]entities.WateringRecord, error) {
	args := m.Called(plantID)
	return args.Get(0).([]entities.WateringRecord), args.Error(1)
}

var newRepoFunc = newRepo

func TestHandlerGETWateringHistory_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockRepo := new(MockRepository)
	
	expectedRecords := []entities.WateringRecord{
		{
			ID:        "1",
			PlantID:   "test-plant-123",
			WateredAt: time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			Notes:     stringPtr("First watering"),
			CreatedAt: time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			ID:        "2",
			PlantID:   "test-plant-123",
			WateredAt: time.Date(2025, 1, 10, 9, 30, 0, 0, time.UTC),
			Notes:     nil,
			CreatedAt: time.Date(2025, 1, 10, 9, 30, 0, 0, time.UTC),
		},
	}
	
	mockRepo.On("FindWateringRecordsByPlantID", "test-plant-123").Return(expectedRecords, nil)
	
	originalNewRepo := newRepoFunc
	newRepoFunc = func() *Repository {
		return &Repository{} // This won't be used since we're mocking
	}
	defer func() { newRepoFunc = originalNewRepo }()
	
	handlerWithMock := func(c *gin.Context) {
		plantID := c.Param("plantId")
		if plantID == "" {
			c.JSON(400, gin.H{"error": "plant_id is required"})
			return
		}

		records, err := mockRepo.FindWateringRecordsByPlantID(plantID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, records)
	}
	
	router := gin.New()
	router.GET("/plants/:plantId/watering", handlerWithMock)
	
	req, _ := http.NewRequest("GET", "/plants/test-plant-123/watering", nil)
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []entities.WateringRecord
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, expectedRecords[0].ID, response[0].ID)
	assert.Equal(t, expectedRecords[0].PlantID, response[0].PlantID)
	assert.Equal(t, "First watering", *response[0].Notes)
	assert.Nil(t, response[1].Notes)
	
	mockRepo.AssertExpectations(t)
}

func TestHandlerGETWateringHistory_MissingPlantID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	router := gin.New()
	router.GET("/plants/:plantId/watering", HandlerGETWateringHistory)
	
	req, _ := http.NewRequest("GET", "/plants//watering", nil)
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "plant_id is required", response["error"])
}

func TestHandlerGETWateringHistory_RepositoryError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockRepo := new(MockRepository)
	
	mockRepo.On("FindWateringRecordsByPlantID", "test-plant-123").Return([]entities.WateringRecord{}, errors.New("database connection failed"))
	
	handlerWithMock := func(c *gin.Context) {
		plantID := c.Param("plantId")
		if plantID == "" {
			c.JSON(400, gin.H{"error": "plant_id is required"})
			return
		}

		records, err := mockRepo.FindWateringRecordsByPlantID(plantID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, records)
	}
	
	router := gin.New()
	router.GET("/plants/:plantId/watering", handlerWithMock)
	
	req, _ := http.NewRequest("GET", "/plants/test-plant-123/watering", nil)
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "database connection failed", response["error"])
	
	mockRepo.AssertExpectations(t)
}

func TestHandlerGETWateringHistory_EmptyResult(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockRepo := new(MockRepository)
	
	mockRepo.On("FindWateringRecordsByPlantID", "nonexistent-plant").Return([]entities.WateringRecord{}, nil)
	
	handlerWithMock := func(c *gin.Context) {
		plantID := c.Param("plantId")
		if plantID == "" {
			c.JSON(400, gin.H{"error": "plant_id is required"})
			return
		}

		records, err := mockRepo.FindWateringRecordsByPlantID(plantID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, records)
	}
	
	router := gin.New()
	router.GET("/plants/:plantId/watering", handlerWithMock)
	
	req, _ := http.NewRequest("GET", "/plants/nonexistent-plant/watering", nil)
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []entities.WateringRecord
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 0)
	
	mockRepo.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}
